package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "encoding/json"
    "crypto/tls"
)



// https://api.coin-swap.net/market/stats/DOGE/BTC

func get_content() {
    tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
    }
    client := &http.Client{Transport: tr}
    // api url that responds with json data
    //url := "https://api.coin-swap.net/market/stats/DOGE/BTC"
    url := "https://api.coin-swap.net/market/summary"

    // Request the url data
    urlResponse, urlError := client.Get(url)

    // If there was an error:
    if urlError != nil {
            fmt.Printf("%s",urlError)
    }

    apiResponse,apiError := ioutil.ReadAll(urlResponse.Body)
    urlResponse.Body.Close() // Close the url request

    if apiError != nil {
        fmt.Printf("%s",apiError)
    }

    var marketStats interface{}
    err := json.Unmarshal(apiResponse, &marketStats)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    // Print json data to screen
    stats_map := marketStats.(map[string]interface{})
    
    fmt.Printf("Results: %v\n", stats_map["symbol"])

    }

func main() {
    get_content()
}