package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type todo struct {
	UserID    int    `json:"userID"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `
<h1> TODO #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>
`

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	const base = "https://jsonplaceholder.typicode.com/"
	resp, err := http.Get(base + r.URL.Path[1:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	var item todo
	err = json.NewDecoder(resp.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create a template with name mine
	tmpl := template.New("mine")

	tmpl.Parse(form)
	tmpl.Execute(w, item)

}
