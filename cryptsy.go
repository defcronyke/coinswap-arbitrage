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
func Get_cryptsy() map[string]float64 {
    tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
    }
    client := &http.Client{Transport: tr}
    // api url that responds with json data
    url := "http://pubapi.cryptsy.com/api.php?method=marketdatav2"

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
    
    stats := marketStats.(map[string]interface{})
    
    //fmt.Printf("%v", stats)
    
	return_data := map[string]float64{}
	return_data = make(map[string]float64)
	
	switch marketStats.(type) {
	    case map[string]interface{}:

			if stats["return"] != nil {
		        
				for market_name, market_data := range stats["return"].(map[string]interface{})["markets"].(map[string]interface{}) {	// For each market.
				    
				    market_name = strings.Replace(market_name, "/", "-", -1)
				    
				    dash_pos := strings.Index(market_name, "-")		// Okay, it looks like Cryptsy does the market names your way Ian...
				    a := market_name[0:dash_pos]
				    b := market_name[dash_pos+1:len(market_name)]
				    market_name = b + "-" + a
				    
				    bid, _ := strconv.ParseFloat(market_data.(map[string]interface{})["buyorders"].([]interface{})[0].(map[string]interface{})["price"].(string), 64)	// WTF?!
				    
				    return_data[market_name] = bid
				    
				}
			}
			
		case nil:
	}
	return return_data
}
