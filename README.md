# Go KRX
Golang으로 KRX 정보를 얻어옵니다 

# How to Install

```shell
go get github.com/dojinkimm/go-krx
```

# How to Use

```go
symbol, err := krx.GetSymbolByCompanyName("삼성전자")

kospi, err := krx.GetKOSPI()

stock, err := krx.GetStockBySymbolAndDate("005930", 3)

topBuyers, err := krx.GetTopBuyersBySymbol("005930")
```



### NOTE:
[KRX 상장지원 서비스](https://kasp.krx.co.kr/contents/02/02010000/ASP02010000.jsp) 홈페이지가 제대로 작동하지 않을시 응답이 정확하지 않을 수 있습니다.