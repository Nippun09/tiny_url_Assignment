package server

import (
	"Infracloud_Assignment/service"
	"log"
	"net/http"
)

func StartUrlShortenServer() {
	http.HandleFunc("/shorten", service.UrlShortenerService)

	lasErr := http.ListenAndServe(":9091", nil)

	if lasErr != nil {
		log.Fatal(lasErr)
	}
}
