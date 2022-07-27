package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
1. Initialize http server.
2. Create route handler

*/
func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	//Create Route handlers
	http.HandleFunc("/welcome", welcomeFunction)
	http.HandleFunc("/profile", profileFunction)

	// Initialize HTTP Server

	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func welcomeFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/welcome" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Welcome Guys!")
}

func profileFunction(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")

	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
	fmt.Fprintf(w, "Phone = %s\n", phone)
	fmt.Fprintf(w, "Address = %s\n", address)
}
