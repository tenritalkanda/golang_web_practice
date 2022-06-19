package handler

import (
	"golang_web/entity"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World Golang Web"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		log.Printf(r.URL.Path)
		return
	}

	data := map[string]interface{}{
		"title":   "Home",
		"content": "Golang Web",
	}

	tmpl, err := template.ParseFiles("views/index.html", "views/layout.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.NotFound(w, r)
		return
	}

}

func ProductHandler(w http.ResponseWriter, r *http.Request) {

	data := []entity.Product{
		{ID: 1, Name: "Product 1", Price: "100", Stocks: 1},
		{ID: 2, Name: "Product 2", Price: "200", Stocks: 10},
		{ID: 3, Name: "Product 3", Price: "300", Stocks: 0},
	}

	tmpl, err := template.ParseFiles("views/product.html", "views/layout.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	// idNumb, err := strconv.Atoi(id)

	// if err != nil || idNumb < 1 {
	// 	http.NotFound(w, r)
	// 	return
	// }
	// // w.Write([]byte(id))

	// fmt.Fprintf(w, "Product page : %d", idNumb)
}

func PostGetHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("GET"))
	case "POST":
		w.Write([]byte("POST"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("views/form.html", "views/layout.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func Formsave(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		name := r.FormValue("name")
		price := r.FormValue("price")
		stocks := r.FormValue("stocks")
		stockInt, err := strconv.Atoi(stocks)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := entity.Product{
			ID:     1,
			Name:   name,
			Price:  price,
			Stocks: stockInt,
		}

		// data := make(map[string]interface{})
		// data["name"] = name
		// data["price"] = price
		// data["stocks"] = stockInt

		tmpl, err := template.ParseFiles("views/result.html", "views/layout.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
