package main

import (
	"fmt"
	"log"
	"net/url"
	"net/http"
)

func generateURL() string{
	u, err := url.Parse("/params")
	if err != nil{
		log.Fatal(err)
	}
	u.Scheme = "http"
	u.Host = "localhost:3000"
	
	q := u.Query()
	q.Set("name", "value")
	u.RawQuery = q.Encode()

	return u.String()
}

func main() {
	url := generateURL()
}

