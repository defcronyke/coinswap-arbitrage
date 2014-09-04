package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "encoding/json"
    "crypto/tls"
)

// returns the golang equivalent of { "market_name": bid_price, ... }
func Get_bittrex() map[string]float64 {
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

		switch marketStats.(type) {
		    case map[string]interface{}:

				for _, v := range markets {		// For each market.
				    market_data := v.(map[string]interface{})
				    
				    if market_data["MarketName"] != nil && market_data["Bid"] != nil {
			    		return_data[market_data["MarketName"].(string)] = market_data["Bid"].(float64)	// Add a market_name -> ask_price pair to the map that we're returning.
			    	}
		    	}
			
			case nil:
				// do nothing
				
		}
	

	return return_data

}
