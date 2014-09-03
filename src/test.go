package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "encoding/json"
    "crypto/tls"
    "./poloniex"
    "./bittrex"
    "./mintpal"
)



// https://api.coin-swap.net/market/stats/DOGE/BTC

func get_content() {	// TODO: Make this return a map[string]float64 where the key is the market name and the value is the top ask price.
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
    		
//    		fmt.Printf("Coin-Swap:\n")
//    		for _,v := range stats {
//    			fmt.Printf("%v\n", v)    
//    		}
    		
    		fmt.Printf("Coin-Swap: %v", stats)
    }
}

func main() {
    get_content()	// Coin-Swap data
    poloniex_data := poloniex.Get_content()
    bittrex_data := bittrex.Get_content()
    mintpal_data := mintpal.Get_content()
    
    fmt.Printf("Poloniex: %v\n", poloniex_data)
    fmt.Printf("Bittrex: %v\n", bittrex_data)
    fmt.Printf("Mintpal: %v\n", mintpal_data)
}