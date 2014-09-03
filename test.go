package main

import (
    "fmt"

)


func main() {
    coinswap_data := Get_coinswap()
    poloniex_data := Get_poloniex()
    bittrex_data := Get_bittrex()
    mintpal_data := Get_mintpal()
    
    _ = Compare(coinswap_data,poloniex_data)
    _ = Compare(coinswap_data,bittrex_data)
    _ = Compare(coinswap_data,mintpal_data)


    
    
/*
    fmt.Printf("Coin-Swap: %v\n", coinswap_data)
    fmt.Printf("Poloniex: %v\n", poloniex_data)
    fmt.Printf("Bittrex: %v\n", bittrex_data)
    fmt.Printf("Mintpal: %v\n", mintpal_data)
*/
}

func Compare(coinswap,exchange map[string]float64) int {
    for market,price := range coinswap {
        if market,ok := exchange[market]; ok {
            fmt.Printf("Matching market: %v\n", market)
            if price < exchange[market] {
                fmt.Printf("Arbitrage available: %v\n", market)
            }
        } else {
            fmt.Printf("No match: %v\n", market)
        }
    }
    return 1;
}