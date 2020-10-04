package main

import "fmt"

func main() {
	symbol, _ := GetSymbolByCompanyName("삼성전자")
	stock, _ := GetStockByDateSymbol(symbol, 10)
	fmt.Println(len(stock))
}
