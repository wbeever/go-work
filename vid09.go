package main

import (
	"fmt"
	"net/http"
)

func indexHandleR(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hey there.</h1>
<p>line one.</p>
<p>line two.</p>
			`)

}

func main() {
	http.HandleFunc("/", indexHandleR)
	http.ListenAndServe(":8000", nil)
}
