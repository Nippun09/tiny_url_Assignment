package service

import (
	"Infracloud_Assignment/db"
	"Infracloud_Assignment/utils"
	"fmt"
	"io/ioutil"
	"net/http"
)

var dbHashcodeToShort map[interface{}]string
var dbShortUrlToLong map[string]string

func init() {
	dbHashcodeToShort = db.GetHashcodeToShortDb()
	dbShortUrlToLong = db.GetShortUrlToLongDb()
}

func UrlShortenerService(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		shortUrlBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Fprintf(responseWriter, "error: failed to read request body")
		}
		if longUrl, isPresent := dbShortUrlToLong[string(shortUrlBytes)]; isPresent {
			fmt.Fprintf(responseWriter, longUrl) //writes long Url to response
		} else {
			fmt.Fprintf(responseWriter, "error: long URL for given short URL doesnt exist")
		}
	case "POST":
		longUrlBytes, err := ioutil.ReadAll(request.Body)
		longUrlString := string(longUrlBytes)
		if err != nil {
			fmt.Fprintf(responseWriter, "error: failed to read request body")
		}
		hashcode := utils.GenerateHash(longUrlString)
		if shorturl, isPresent := dbHashcodeToShort[hashcode]; isPresent {
			fmt.Fprintf(responseWriter, shorturl) //writes short url to response
		} else {
			hashcode := utils.GenerateHash(string(longUrlBytes))
			shorturl := utils.Shorten(longUrlString)
			dbHashcodeToShort[hashcode] = shorturl
			dbShortUrlToLong[shorturl] = longUrlString
			fmt.Fprintf(responseWriter, shorturl) //writes short url to response
		}
	default:
		fmt.Fprintf(responseWriter, "error: Invalid request method, POST methods are supported.")
	}
}
