package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

//SiteMapIndex looks for the "sitemap" tag??
type SiteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

//News sub tag of linked url?
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

//NewsMap yeah.. idk?
type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
	var s SiteMapIndex
	var n News
	newsMap := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for idx := range n.Titles {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	for idx, data := range newsMap {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)

	}
}
