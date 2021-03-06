package tests

import (
	"Infracloud_Assignment/utils"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func testShortenUrl(longUrl string) (string, bool) {

	requestBody := bytes.NewBuffer([]byte(longUrl))
	response, err := http.Post("http://localhost:8080", "application/text", requestBody)
	if err != nil {
		log.Fatal("failed to shorten url :", err)
		return "", false
	}
	defer response.Body.Close()

	responseBody, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal("failed to read from response :", readErr)
		return "", false
	}

	shortUrl := utils.Shorten(longUrl)
	log.Printf("short url is: %v", shortUrl)
	log.Printf("response body is: %v", responseBody)
	if shortUrl == string(responseBody) {
		return shortUrl, true
	}
	return "", false
}
