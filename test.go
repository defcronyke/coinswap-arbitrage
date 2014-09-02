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
        marketid string
        symbol string
        exchange string
        lastprice string
        dayvolume string
        dayhigh string
        daylow string
        ask string
        bid string
        openorders string
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
            fmt.Printf("%s",urlError)
    }

    //apiDecoder := json.NewDecoder(strings.NewReader(apiResponse))
    var jsonData []marketStats

    apiDecoder := json.Unmarshal(apiResponse, &jsonData)
    fmt.Printf("Results: %v\n", apiDecoder)

    }

func main() {
    get_content()
}