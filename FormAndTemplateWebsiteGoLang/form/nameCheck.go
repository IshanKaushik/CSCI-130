package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/results", myResultDisplay)
	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	indexPage, err := ioutil.ReadFile("pages/index.html")
	if err != nil {
		fmt.Fprint(w, "404 Not Found")
	}
	fmt.Fprint(w, string(indexPage))
}

var output, _ = ioutil.ReadFile("pages/checkResult.html")
var ouputTemplate = template.Must(template.New("output").Parse(string(output)))

func myResultDisplay(w http.ResponseWriter, r *http.Request) {
	strEntered := r.FormValue("name")
	if strEntered == "ISHAN KAUSHIK" {
		err := ouputTemplate.Execute(w, "WOW , you remembring your name.")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		err := ouputTemplate.Execute(w, "Try hard !!!")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
