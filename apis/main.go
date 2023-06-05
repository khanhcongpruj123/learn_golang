package main

import (
	"net/http"
)

func main() {
	// http.Handle("/", TaskRouter())
	http.Handle("/", AuthRouter())
	http.ListenAndServe(":8080", nil)
}
