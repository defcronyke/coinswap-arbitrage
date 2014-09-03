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
    
	var marketStats interface{}


    // Request the url data
    urlResponse, urlError := client.Get(url)

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

    err := json.Unmarshal(apiResponse, &marketStats)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    // Print json data to screen
    switch marketStats.(type) {
        case map[string]interface{}:	// If json response type is a map with string keys.
        	stats := marketStats.(map[string]interface{})
    		fmt.Printf("Results: %v\n", stats["marketid"])
    	case []interface{}:				// if type is an array.
    		stats := marketStats.([]interface{})
    		
    		fmt.Printf("Results:\n")
    		for _,v := range stats {
    			fmt.Printf("%v\n", v)    
    		}
    		
    }


}

func main() {
    get_content()
}