package main

import (
    "fmt"

)


func main() {
    coinswap_data := Get_coinswap()
    //poloniex_data := poloniex.Get_content()
    //bittrex_data := bittrex.Get_content()
    //mintpal_data := mintpal.Get_content()

    fmt.Printf("Coin-Swap: %v\n", coinswap_data)
    //fmt.Printf("Poloniex: %v\n", poloniex_data)
    //fmt.Printf("Bittrex: %v\n", bittrex_data)
    //fmt.Printf("Mintpal: %v\n", mintpal_data)
}