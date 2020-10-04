package main

type Stock struct {
	quote         string
	upDown        string
	dayOverDay    string
	previousClose string
	volume        string
	tradingValue  string
	marketValue   string
	low           string
	high          string
	low52Week     string
	high52week    string
	upperLimit    string
	lowerLimit    string
	listedShares  string
	faceValue     string
	per           string
}

func GetStockBySymbol(symbol string) (*Stock, error) {
	s, err := getStockInfoBySymbol(symbol)
	if err != nil {
		return nil, err
	}

	return &Stock{
		quote:         s.TBLStockInfo.CurJuka,
		upDown:        s.TBLStockInfo.DungRak,
		dayOverDay:    s.TBLStockInfo.Debi,
		previousClose: s.TBLStockInfo.PrevJuka,
		volume:        s.TBLStockInfo.Volume,
		tradingValue:  s.TBLStockInfo.Money,
		marketValue:   s.TBLStockInfo.StartJuka,
		low:           s.TBLStockInfo.LowJuka,
		high:          s.TBLStockInfo.HighJuka,
		low52Week:     s.TBLStockInfo.Low52,
		high52week:    s.TBLStockInfo.High52,
		upperLimit:    s.TBLStockInfo.UpJuka,
		lowerLimit:    s.TBLStockInfo.DownJuka,
		listedShares:  s.TBLStockInfo.Amount,
		faceValue:     s.TBLStockInfo.FaceJuka,
		per:           s.TBLStockInfo.Per,
	}, nil
}
