package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

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

type newsPage struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, go is neat!")
}

func newsRoutine(c chan newsPage, Location string) {
	defer wg.Done()
	var n newsPage
	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()
	c <- n
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s siteMapIndex
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	newsVar := make(map[string]newsMap)
	resp.Body.Close()
	queue := make(chan newsPage, 30)

	for _, Location := range s.Locations {
		wg.Add(1)
		go newsRoutine(queue, Location)
	}
	wg.Wait()
	close(queue)
	for elem := range queue {
		for idx := range elem.Titles {
			newsVar[elem.Titles[idx]] = newsMap{elem.Keywords[idx], elem.Locations[idx]}
		}
	}

	p := newsAggPage{Title: "Amazing New Aggregator", News: newsVar}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
