/*
first try of golang
*/

package main

import (
	"fmt"
	"net/http"
)

func landing_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Trying for Go dev")
}

func discription_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "DISCRIPTION")
}

func handleRequest() {
	http.HandleFunc("/", landing_page)
	http.HandleFunc("/discription", discription_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
