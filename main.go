package main

import (
	"net/http"
	"url-shortener/handlers"
	"url-shortener/store"
)

func main() {
	urlStore := store.NewURLStore()
	urlHandler := handlers.NewURLHandler(urlStore)

	http.HandleFunc("/shorten", urlHandler.ShortenURL)
	http.HandleFunc("/topdomains", urlHandler.TopDomains)
	http.HandleFunc("/", urlHandler.Redirect)

	http.ListenAndServe(":8080", nil)
}
