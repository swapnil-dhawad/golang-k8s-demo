package main

import (
	"log"
	"net/http"
)


func welcomePage(w http.ResponseWriter, r *http.Request) {
	// Render the course html page
	http.ServeFile(w, r, "static/welcome.html")
}



func main() {

	http.HandleFunc("/", welcomePage)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
