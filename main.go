package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"ascii-web/funcs"
)

var (
	tpl     *template.Template
	port    = "8080"
	text    string
	banner  string
	err     error
	ErrType string
	inputs  = []string{"test", "Hello\r\n\r\nThere!", "\r\n\r\n\r\n\r\nI'm after 4 new lines", "I'm before 2 new lines\r\n\r\n", "ABCDEFghij123@", "123,;:!456", "Special Characters\r\n{:>|}{}", ">>NonPrintable:é<<"}
	values  struct {
		Btn     string
		Output  string
		Outputs []string
	}
)

// a func for return the statut error code
func ErrorHandle(myStrStatus string) int {
	if myStrStatus == "internal" {
		return http.StatusInternalServerError
	} else if myStrStatus == "method" {
		return http.StatusMethodNotAllowed
	} else if myStrStatus == "request" {
		return http.StatusBadRequest
	} else if myStrStatus == "notfound" {
		return http.StatusNotFound
	}

	return 0
}

func ExecError(ErrType string, w http.ResponseWriter) {
	tpl, _ = template.ParseGlob("static/templates/error.html")
	w.WriteHeader(ErrorHandle(ErrType))
	tpl.ExecuteTemplate(w, "error.html", http.StatusText(ErrorHandle(ErrType)))
}

// a func return the file html who's not exist in the templates directory
func InternalErr() string {
	MySlice := []string{"error", "output", "index"}
	for _, v := range MySlice {
		tpl, err = template.ParseGlob("static/templates/" + v + ".html")
		if err != nil {
			return v
		}
	}
	return ""
}

func IndexHanlde(w http.ResponseWriter, r *http.Request) {
	// ---------------------check if html files is exist-------------------------
	if InternalErr() == "error" {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if InternalErr() == "index" || InternalErr() == "output" {
		ExecError("internal", w)
		return
	}
	// --------------------------------------------------------------------------

	// ----------------------check the banners if exist----------------------------
	if !funcs.BannerExist() {
		ExecError("internal", w)
		return
	}

	// -----------------------------handling path---------------------------------
	if r.URL.Path != "/" {
		ExecError("notfound", w)
		return
	}

	// ------------------------check the method request-----------------------------
	if r.Method != http.MethodGet {
		ExecError("method", w)
		return
	}

	// ---------------if all good the index template execute---------------------
	tpl, _ = template.ParseGlob("static/templates/index.html")
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func OutputHandle(w http.ResponseWriter, r *http.Request) {

	// ---------------------check html files is exist----------------------------
	if InternalErr() == "error" {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if InternalErr() == "index" || InternalErr() == "output" {
		ExecError("internal", w)
		return
	}
	// --------------------------------------------------------------------------

	// ----------------------check the banners if exist----------------------------
	if !funcs.BannerExist() {
		ExecError("internal", w)
		return
	}

	// ------------------------check the method request-----------------------------
	if r.Method != http.MethodPost {
		ExecError("method", w)
		return
	}

	data := values

	data.Outputs = []string{}            // for store the output of liste of tests
	data.Output = ""                     // for store the output of user input
	banner = r.FormValue("banner")       // store the banner
	text = r.FormValue("textName")       // store the input text
	data.Btn = r.FormValue("ButtonName") // store which button the user click

	if data.Btn == "submit" {

		data.Output = funcs.Print(text, banner)
		// ------------------check if the user enter the text and banner--------------------
		if (len(banner) != 0 && len(text) != 0) && data.Output != "err" && len(text) <= 1000 {
			tpl, _ = template.ParseGlob("static/templates/output.html")
			tpl.ExecuteTemplate(w, "output.html", data)
			return

		} else {
			ExecError("request", w)
			return

		}

	} else if data.Btn == "test" {

		if len(banner) == 0 || (banner != "standard" && banner != "thinkertoy" && banner != "shadow") {
			// ------if the user don't chose any banner the test run with standard banner------
			banner = "standard"
		}

		for _, value := range inputs {
			data.Outputs = append(data.Outputs, funcs.Print(value, banner))
		}

		tpl, _ = template.ParseGlob("static/templates/output.html")
		tpl.ExecuteTemplate(w, "output.html", data)

	} else {
		ExecError("request", w)
		return

	}
}

func StaticHandle(w http.ResponseWriter, r *http.Request) {
	ExecError("notfound", w)
}

func main() {
	fmt.Println("http://localhost:" + port)
	http.HandleFunc("/", IndexHanlde)
	http.HandleFunc("/static", StaticHandle)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/ascii-art", OutputHandle)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
