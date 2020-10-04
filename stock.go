package main

type Stock struct {
	quote            string
	upDown           string
	comparedToPrev   string
	prevEndPrice     string
	volume           string
	transactionPrice string
	marketPrice      string
	lowPrice         string
	highPrice        string
	low52WeekPrice   string
	high52WeekPrice  string
	upperLimit       string
	lowerLimit       string
	listedShares     string
	faceValue        string
	per              string
}

func GetStockBySymbol(symbol string) (*Stock, error) {
	s, err := getStockInfoBySymbol(symbol)
	if err != nil {
		return nil, err
	}

	return &Stock{
		quote:            s.TBLStockInfo.CurJuka,
		upDown:           s.TBLStockInfo.DungRak,
		comparedToPrev:   s.TBLStockInfo.Debi,
		prevEndPrice:     s.TBLStockInfo.PrevJuka,
		volume:           s.TBLStockInfo.Volume,
		transactionPrice: s.TBLStockInfo.Money,
		marketPrice:      s.TBLStockInfo.StartJuka,
		lowPrice:         s.TBLStockInfo.LowJuka,
		highPrice:        s.TBLStockInfo.HighJuka,
		low52WeekPrice:   s.TBLStockInfo.Low52,
		high52WeekPrice:  s.TBLStockInfo.High52,
		upperLimit:       s.TBLStockInfo.UpJuka,
		lowerLimit:       s.TBLStockInfo.DownJuka,
		listedShares:     s.TBLStockInfo.Amount,
		faceValue:        s.TBLStockInfo.FaceJuka,
		per:              s.TBLStockInfo.Per,
	}, nil
}
