package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "encoding/json"
    "crypto/tls"
    "strconv"
    "strings"
)

// returns the golang equivalent of { "market_name": bid_price, ... }
func Get_poloniex() map[string]float64 {
    tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
    }
    client := &http.Client{Transport: tr}
    // api url that responds with json data
    url := "https://poloniex.com/public?command=returnTicker"

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
	return_data := map[string]float64{}
	return_data = make(map[string]float64)
	    
	    switch marketStats.(type) {
		    case map[string]interface{}:
	    
				for k, v := range stats_map {	// For each market.
				    market_data := v.(map[string]interface{})
				    
				    k = strings.Replace(k, "_", "-", -1)	// Poloniex uses BTC_DOGE format for market names, but we want BTC-DOGE.
				    
				    if market_data["highestBid"] != nil {
				    	return_data[k], _ = strconv.ParseFloat(market_data["highestBid"].(string), 64)	// Poloniex prices are strings, so we need to make them into floats.
			    	}
				}
				
			case nil:
				// do nothing
		}
	
	return return_data
}

