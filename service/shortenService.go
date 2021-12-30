package service

import (
	"Infracloud_Assignment/db"
	"Infracloud_Assignment/utils"
	"fmt"
	"io/ioutil"
	"net/http"
)

var dbHashcodeToShort map[uint32]string
var dbShortUrlToLong map[string]string

func init() {
	dbHashcodeToShort = db.GetHashcodeToShortDb()
	dbShortUrlToLong = db.GetShortUrlToLongDb()
}

func UrlShortenerService(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	// case "GET":
	// 	shortUrlBytes, err := ioutil.ReadAll(request.Body)
	// 	if err != nil {
	// 		fmt.Fprintf(responseWriter, "error: failed to read request body")
	// 	}
	// 	if longUrl, isPresent := dbShortUrlToLong[string(shortUrlBytes)]; isPresent {
	// 		fmt.Fprintf(responseWriter, longUrl) //writes long Url to response
	// 	} else {
	// 		fmt.Fprintf(responseWriter, "error: long URL for given short URL doesnt exist")
	// 	}
	case "POST":
		longUrlBytes, err := ioutil.ReadAll(request.Body)
		longUrlString := string(longUrlBytes)
		fmt.Printf("long url string :%s\n", longUrlString)
		if err != nil {
			fmt.Fprintf(responseWriter, "error: failed to read request body")
		}
		fmt.Println("before generate hash of longurl string")
		hashcode := utils.GenerateHash(longUrlString)
		fmt.Println("after generate hash of longurl string")
		if shorturl, isPresent := dbHashcodeToShort[hashcode]; isPresent {
			fmt.Fprintf(responseWriter, shorturl) //writes short url to response
		} else {
			fmt.Println("before generate hash of longurl string inside else")
			hashcode := utils.GenerateHash(string(longUrlBytes))
			fmt.Println("after generate hash of longurl string inside else")
			shorturl := utils.Shorten(longUrlString)
			fmt.Println("after shortening of long url string")
			dbHashcodeToShort[hashcode] = shorturl
			dbShortUrlToLong[shorturl] = longUrlString
			fmt.Fprintf(responseWriter, shorturl) //writes short url to response
		}
	default:
		fmt.Fprintf(responseWriter, "error: Invalid request method, POST methods are supported.")
	}
}
