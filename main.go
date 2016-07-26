package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
)

func EventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
	w.Write([]byte("Hello World!"))	
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", EventHandler)

	fmt.Println("Server is Listening on 10110")

	log.Fatal(http.ListenAndServe(":10110", r))
}
