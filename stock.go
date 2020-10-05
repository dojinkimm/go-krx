package krx

type Stock struct {
	Quote         string
	UpDown        UpDown
	DayOverDay    string
	PreviousClose string
	Volume        string
	TradingValue  string
	StartValue    string
	Low           string
	High          string
	Low52Week     string
	High52week    string
	UpperLimit    string
	LowerLimit    string
	ListedShares  string
	FaceValue     string
	PER           string
}

func GetStockBySymbol(symbol string) (*Stock, error) {
	s, err := getStockInfoBySymbol(symbol)
	if err != nil {
		return nil, err
	}

	return &Stock{
		Quote:         s.TBLStockInfo.CurJuka,
		UpDown:        castToUpDown(s.TBLStockInfo.DungRak),
		DayOverDay:    s.TBLStockInfo.Debi,
		PreviousClose: s.TBLStockInfo.PrevJuka,
		Volume:        s.TBLStockInfo.Volume,
		TradingValue:  s.TBLStockInfo.Money,
		StartValue:    s.TBLStockInfo.StartJuka,
		Low:           s.TBLStockInfo.LowJuka,
		High:          s.TBLStockInfo.HighJuka,
		Low52Week:     s.TBLStockInfo.Low52,
		High52week:    s.TBLStockInfo.High52,
		UpperLimit:    s.TBLStockInfo.UpJuka,
		LowerLimit:    s.TBLStockInfo.DownJuka,
		ListedShares:  s.TBLStockInfo.Amount,
		FaceValue:     s.TBLStockInfo.FaceJuka,
		PER:           s.TBLStockInfo.Per,
	}, nil
}
