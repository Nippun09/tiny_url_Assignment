# Infracloud_Assignment

This is URL shortener project.

Go version 1.16.7

This project uses hashing approach for generating corresponding short URL from given long URL.

Start folder contains main function to start server.

environment setup for running this project: set GO111MODULE=off.

step 1: IDE terminal run following command "go env -w GO111MODULE=off"
step 2: run main.go present in Infracloud_Assignment/start folder using command go run <from/current_directory/path/to/main.go>
step 3:use postman to test the URL
    3.a: use POST method to hit following url "http://localhost:8080/shorten"
    3.b: in body section select 'raw' and keep text option as 'text'
    3.c: past the url to be shortened in the body and click on send button.
    3.d: you will get response as corresponding short url.