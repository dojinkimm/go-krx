package main

import (
	"fmt"

	"github.com/dojinkimm/go-krx"
)

func main() {
	kospi, err := krx.GetKOSPI()
	fmt.Println(kospi, err)

	symbol, err := krx.GetSymbolByCompanyName("삼성전자")
	fmt.Println(symbol, err)
}
