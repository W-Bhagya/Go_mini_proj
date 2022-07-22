package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(w,"ParseForm() err : %v",err)
		return
	}

	fmt.Println(w,"POST request successful")
	name := r.FormValue("name")
	email := r.FormValue("email")
	empid := r.FormValue("empid")
	address := r.FormValue("address")


	fmt.Fprintf(w, "Name = %s\n",name)
	fmt.Fprintf(w, "Email = %s\n",email)
	fmt.Fprintf(w, "EmpId = %s\n",empid)
	fmt.Fprintf(w, "Address = %s\n",address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"Hello....!")
}


func main() {

	// checks static directory
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	
	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
}