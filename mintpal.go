package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "encoding/json"
    "crypto/tls"
    "strconv"
//    "strings"
)

// returns the golang equivalent of { "market_name": ask_price, ... }
func Get_mintpal() map[string]float64 {
    tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
    }
    client := &http.Client{Transport: tr}
    // api url that responds with json data
    url := "https://api.mintpal.com/v1/market/summary/"

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
    
    stats_array := marketStats.([]interface{})
    
    //fmt.Printf("%v", stats_array)
    
	return_data := map[string]float64{}
	return_data = make(map[string]float64)

	for _, v := range stats_array {	// For each market.
	    market_data := v.(map[string]interface{})
	    return_data[market_data["exchange"].(string)+"-"+market_data["code"].(string)], _ = strconv.ParseFloat(market_data["top_ask"].(string), 64)	// Mintpal prices are strings, so we make them into floats.
	}
	
	return return_data
}

//func main() {
//    mintpal_data := get_content()
//    
//    for k, v := range mintpal_data {
//        fmt.Printf("%v: %.8f\n", k, v)
//    }
//    
//}