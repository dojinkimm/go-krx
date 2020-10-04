package main

import "fmt"

func main() {
	symbol, _ := GetSymbolByCompanyName("삼성전자")
	sellers, _ := GetTopSellersBySymbol(symbol)
	for _, s := range sellers {
		fmt.Println(s.name, s.volume)
	}
}
