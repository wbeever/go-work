package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type newsAggPage struct {
	Title string
	News  string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := newsAggPage{Title: "Amazing New Aggregator", News: "some news"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, go is neat!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
