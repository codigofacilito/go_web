package main

import (
    "fmt"
    "log"
    "net/url"
    "net/http"
    "io/ioutil"
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
    req, err := http.NewRequest("GET", url, nil)
    if err != nil{
        panic(err)
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}



