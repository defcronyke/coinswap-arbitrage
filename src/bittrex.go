package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "encoding/json"
    "crypto/tls"
)

// returns the golang equivalent of { "market_name": ask_price, ... }
func get_content() map[string]float64 {
    tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
    }
    client := &http.Client{Transport: tr}
    // api url that responds with json data
    url := "https://bittrex.com/api/v1.1/public/getmarketsummaries"

	var marketStats interface{}

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

    err := json.Unmarshal(apiResponse, &marketStats)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    stats_map := marketStats.(map[string]interface{})
	markets := stats_map["result"].([]interface{})

	return_data := map[string]float64{}
	return_data = make(map[string]float64)

	for _, v := range markets {
	    market_data := v.(map[string]interface{})
	    return_data[market_data["MarketName"].(string)] = market_data["Ask"].(float64)
	}
	
	return return_data
}

func main() {
    bittrex_data := get_content()
    
    for k, v := range bittrex_data {
        fmt.Printf("%v: %.8f\n", k, v)
    }
    
}