package main

import "net/http"

func main() {
	http.Get("http://localhost:9200/")
}