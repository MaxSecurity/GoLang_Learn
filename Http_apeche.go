package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	_ = http.ListenAndServe(":2333", nil)
}

//Http_apeche
