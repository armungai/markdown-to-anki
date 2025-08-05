package main // declares a standalone exec. -> main entry point

import (
	"html/template" // -> for parsing and rendering HTML templates safely (prevents XSS)
	"log" // prints logs in the console
	"net/http" // provides http server functionality
)

/*
	Function named handler which handles HTTP requests
	PARAM: w http.ResponseWriter -> writes HTTP responses back to the client
	PARAM: *http.Request -> represents incoming HTTP request
*/

func handler(w http.ResponseWriter, r *http.Request) {
	// Parses the HTML template file located at said address
	// template.Must is a helper that panics if error is non-nil -> catches startup problems
	tmpl := template.Must(template.ParseFiles("templates/index.html"))


	// creates an anonymous struct with Title, Heading and Message
	data := struct {
		Title string
		Heading string
	} {
		Title: "Markdown to Anki",
		Heading: "Markdown to Anki",

	}

	// applies the template with the provided data
	// results are written to the http.Response writer so the browser reveives the rendered html
	tmpl.Execute(w, data)
}

// entry point
func main() {
	// Serve files from the "static" directory
	http.HandleFunc("/", handler)

	log.Println(("Server started on http://localhost:8080"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
