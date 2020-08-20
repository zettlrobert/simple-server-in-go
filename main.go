package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// create a struct that holds information to be displayed in the HTML file
type Welcome struct {
	Name string
	Time string
}

//Go application entrypoint
func main() {
	//Instantiate a Welcome struct object and pass in some random information.
	//name of the user as a query parameter from the URL
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	//we need to tell go where exactly the html file lies. go needs to parse the html file
	//a relative path is passed. it is wrapped in a call to template.Must() which handles any errors and halts if there are fatal errors.
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) // go looks in hte relative staticdirectory first, then matches it
			//url of our choice as shown in http.Handle("/static/").
			//this url is twhat we need when referencing css files
			//once the server begins our html cde would therefore be <link rel="stylesheet" href="/static/stylesheet/...">
			//it is important to note the final url can be whatever we like, so long as we are consistent.

	//takes the url path and a function that takes in a response writer, and a http request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//take name from the url query e.g ?name=Robert, will set welcome.Name = Robert.
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		//if errors show an internal server error message
		//pass welcome struct to the welcome-template.html file.
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Serving on PORT 8000")
	//start the web server, set the port. without path  it assumes localhost
	//print errors from starting the websserver using fmg
	fmt.Println(http.ListenAndServe(":8080", nil))
}
