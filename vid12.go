package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

//SiteMapIndex looks for the "sitemap" tag??
type SiteMapIndex struct {
	Locations []Location `xml:"sitemap"`
}

/*
[5]int == array
[]int == slice
*/

//Location looks for the "loc" tag??
type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}
}
