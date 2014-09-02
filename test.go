package main

import (
    "net/http"
    //"io"
    "io/ioutil"
    "fmt"
    "encoding/json"
    //"log"
    //"strings"
)

// https://api.coin-swap.net/market/stats/DOGE/BTC

func get_content() {
    // api url that responds with json data
    url := "https://api.coin-swap.net/market/stats/DOGE/BTC"


    type marketStats struct {
        Marketid string `json:"marketid"`
        Symbol string `json:"symbol"`
        Exchange string `json:"exchange"`
        Lastprice string `json:"lastprice"`
        Dayvolume string `json:"dayvolume"`
        Dayhigh string `json:"dayhigh"`
        Daylow string `json:"daylow"`
        Ask string `json:"ask"`
        Bid string `json:"bid"`
        Openorders string `json:"openorders"`
    }


    // Request the url data
    urlResponse, urlError := http.Get(url)

    // If there was an error:
    if urlError != nil {
            fmt.Printf("%s",urlError)
    }

    // 
    apiResponse,apiError := ioutil.ReadAll(urlResponse.Body)
    urlResponse.Body.Close() // Close the url request

    if apiError != nil {
        fmt.Printf("%s",apiError)
    }

    var jsonData marketStats
    err := json.Unmarshal(apiResponse, &jsonData)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    // Print json data to screen
    fmt.Printf("Results: %v\n", jsonData.Marketid)

    }

func main() {
    get_content()
}