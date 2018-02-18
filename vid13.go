package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

//SiteMapIndex looks for the "sitemap" tag??
type SiteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

//News sub tag of linked url
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	var s SiteMapIndex
	var n News
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
	}
}
