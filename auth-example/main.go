package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Product struct {
	Id int
	Name string
	Slug string
	Description string
}

var products = []Product{
	{1, "World of Authcraft", "world-of-authcraft", "Battle bugs and protect yourself from invaders while you explore a scary world with no security"},
	{2, "Ocean Explorer", "ocean-explorer", "Explore the depths of the sea in this one of a kind underwater experience"},
	{3,"Dinosaur Park","dinosaur-park", "Go back 65 million years in the past and ride a T-Rex"},
	{4,"Cars VR","cars-vr", "Get behind the wheel of the fastest cars in the world."},
	{5,"Robin Hood","robin-hood", "Pick up the bow and arrow and master the art of archery"},
	{6,"Real World VR", "real-world-vr", "Explore the seven wonders of the world in VR"},
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.Handle("/status", StatusHandler).Methods(http.MethodGet)
	r.Handle("/products", ProductsHandler).Methods(http.MethodGet)
	r.Handle("/products/{slug}/feedback", AddFeedbackHandler).Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented"))
})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(products)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(payload))
})

var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	w.Header().Set("content-type", "application/json")
	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write(payload)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
})