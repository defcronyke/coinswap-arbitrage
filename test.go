package main

import (
    "fmt"

)


func main() {
    coinswap_data := Get_coinswap()
    poloniex_data := Get_poloniex()
    bittrex_data := Get_bittrex()
    mintpal_data := Get_mintpal()

    fmt.Printf("Coin-Swap: %v\n", coinswap_data)
    fmt.Printf("Poloniex: %v\n", poloniex_data)
    fmt.Printf("Bittrex: %v\n", bittrex_data)
    fmt.Printf("Mintpal: %v\n", mintpal_data)
}