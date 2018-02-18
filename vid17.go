package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type newsMap struct {
	Keyword  string
	Location string
}

type newsAggPage struct {
	Title string
	News  map[string]newsMap
}

type siteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

//News sub tag of linked url?
type newsPage struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, go is neat!")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s siteMapIndex
	var n newsPage
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	news_map := make(map[string]newsMap)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx := range n.Titles {
			news_map[n.Titles[idx]] = newsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	p := newsAggPage{Title: "Amazing New Aggregator", News: news_map}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
