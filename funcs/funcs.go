package funcs

import (
	"os"
	"strings"
)

func BannerExist() bool {
	MySlice := []string{"shadow", "standard", "thinkertoy"}
	for _, v := range MySlice {
		_, err := os.Stat("banner/" + v + ".txt")
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func IsEmpty(str string) bool {
	return len(str) == 0
}

func ErrInput(input string) bool { // for handling inputs error (if have any unicode character return error)

	var res string

	for _, ch := range input {
		if ch == 10 {
			res += "\\n"
		} else {
			res += string(ch)
		}
	}

	return false
}

func IsFullNewLine(slice []string) bool { // for handling just input with new lines
	for _, ch := range slice {
		if ch != "" {
			return false
		}
	}
	return true
}

func Print(input string, style string) string {
	databyte, err := os.ReadFile("banner/" + style + ".txt")
	if err != nil {
		return "err"
	}

	content := string(databyte)
	content = strings.ReplaceAll(content, "\r", "") // for handle the error into the thinkertoy file
	letter := strings.Split(content[1:], "\n\n")    // split characters in the standard file
	res := ""

	slice := strings.Split(input, "\r\n") // split the input by new line for print index by index

	for g := 0; g < len(slice); g++ {
		if IsEmpty(slice[g]) { // if the index is empty return new line

			res += "\n"
			if g == len(slice)-1 {
				if IsFullNewLine(slice) {
					res = res[:len(res)-1]
				}

			}

		} else {
			for i := 0; i < 8; i++ { // print word by word and line by line
				for k := 0; k < len(slice[g]); k++ { // print line by line of the words
					if byte(slice[g][k]) < 32 || byte(slice[g][k]) > 126 {
						continue
					}
					word := letter[byte(slice[g][k])-32] // target the character in the standard file
					row := strings.Split(word, "\n")

					for j := 0; j < len(row[i]); j++ { // print line by line

						if row[i][j] != '\n' {
							res += string(row[i][j])
						} else {
							break // stop on the last line of each character
						}
					}
				}

				res += "\n"

			}

		}

	}

	return res
}
