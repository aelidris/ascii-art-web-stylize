# Ascii-Art-Web

Ascii-Art-Web is a web-based application that allows users to convert text into ascii art using different banners. The project consists of a web server that provides a graphical user interface (GUI) for interacting with the ascii-art service.

## Description

The application allows users to:
- Enter text and choose a banner to convert the text into ascii art.
- View the generated ascii art on the webpage.
- Use predefined test inputs to generate ascii art for multiple inputs at once.

## Authors

- [MOHAMMED MIHIT]
- [ABDELFATTAH ELIDRISSI]

## Usage

### Prerequisites

- Go ( go 1.22.3 version or later) installed on your system.
- A web browser to access the application.

### How to Run

1. **Clone the Repository:**
    ```sh
    git clone https://learn.zone01oujda.ma/git/aelidris/ascii-art-web-stylize
    cd ascii-art-web-stylize
    ```

2. **Install Dependencies:**
    Ensure all necessary dependencies are installed. This project relies on the standard library, so no additional dependencies are required.

3. **Run the Server:**
    ```sh
    go run main.go
    ```

4. **Access the Application:**
    Open your web browser and go to [http://localhost:8080/](http://localhost:8080/).

## Implementation Details

### Project Structure

The project consists of the following main components:

- **Templates**: HTML files located in the `static/templates` directory.
- **Main Go File**: The `main.go` file which sets up the web server and handles HTTP requests.

### Endpoints

1. **GET /home**: 
    - Sends the main HTML page where users can input text and select banners.
    - HTTP Status Codes: 
      - `200 OK` if the page loads successfully.
      - `500 Internal Server Error` if there is an issue loading the page.

2. **POST /output**: 
    - Handles the submission of text and banner selection, generating the ascii art.
    - HTTP Status Codes: 
      - `200 OK` if the ascii art is generated successfully.
      - `400 Bad Request` if the input data is invalid or missing.
      - `500 Internal Server Error` if there is an issue generating the ascii art.

### Error Handling

The application includes appropriate error handling to return meaningful HTTP status codes:

- **200 OK**: Request completed successfully.
- **404 Not Found**: Resource not found.
- **405 Method Not Allowed**: This status code indicates that the method specified in the request is not allowed for the resource identified by the request URL.
- **400 Bad Request**: Incorrect requests due to invalid input data.
- **500 Internal Server Error**: Unhandled errors such as server crashes or unexpected conditions.
