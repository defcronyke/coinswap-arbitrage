package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "encoding/json"
    "crypto/tls"
    "strconv"
)

// returns the golang equivalent of { "market_name": ask_price, ... }
func Get_coinswap() map[string]float64 {		// Has to start with a capital leter because we are exporting it.
    tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
    }
    client := &http.Client{Transport: tr}
    // api url that responds with json data
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

    var marketStats interface{} // Create an interface{} type variable to store the api json response
    return_data := map[string]float64{} // Create a map variable that will be used to return our data
    return_data = make(map[string]float64) // Initialize our map

    err := json.Unmarshal(apiResponse, &marketStats)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }

    switch marketStats.(type) {
        case map[string]interface{}:	// If json response type is a map with string keys.
        	stats := marketStats.(map[string]interface{})
                marketname := stats["exchange"].(string) + "-" + stats["symbol"].(string)
                return_data[marketname], _ = strconv.ParseFloat(stats["ask"].(string), 64)	// Add a market_name -> ask_price pair to the map that we're returning.
    	case []interface{}:             // if type is an array.
    		rows := marketStats.([]interface{})
    		for _,v := range rows {
                    stats := v.(map[string]interface{})
     
                    if stats["exchange"] != nil && stats["symbol"] != nil && stats["ask"] != nil {
                    	marketname := stats["exchange"].(string) + "-" + stats["symbol"].(string)
                    	return_data[marketname], _ = strconv.ParseFloat(stats["ask"].(string), 64)	// Add a market_name -> ask_price pair to the map that we're returning.
                   	}
    		}
    }
    return return_data
}
