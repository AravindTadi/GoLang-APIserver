package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Aravind")
	fmt.Println("Endpoint hit")
}
func main() {
	fmt.Println("j")
	http.HandleFunc("/foo", homepage)
	http.ListenAndServe("localhost:10000", nil)

}
