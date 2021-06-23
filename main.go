package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/product", productHandler)
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("masuk profile"))
	})

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World Golang Web"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		log.Printf(r.URL.Path)
		return
	}
	w.Write([]byte("Welcome Home"))
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte(id))

	fmt.Fprintf(w, "Product page : %d", idNumb)
}
