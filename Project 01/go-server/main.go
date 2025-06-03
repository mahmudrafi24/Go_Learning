package main 

import (
	"fmt"
	"log"
	"net/http"
)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello, World!")
}
func formHaldler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	name := r.FormValue("name")
	email := r.FormValue("email")
	if name == "" || email == "" { 
	}
	fmt.Fprintf(w, "Hello, %s!", name) 
}
func main(){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form",formHaldler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server is running on port 8080")
	if err:= http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}