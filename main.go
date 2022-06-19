package main

import (
	"golang_web/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	profileHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("masuk coy"))
	}

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/profile", profileHandler)
	mux.HandleFunc("/postget", handler.PostGetHandler)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/formsave", handler.Formsave)

	assetPath := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", assetPath))

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
