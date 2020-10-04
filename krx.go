package main

type KOSPI struct {
	kospi           string
	kospiUpDown     string
	kospiDayOverDay string
}

type KOSPI200 struct {
	kospi200           string
	kospi200UpDown     string
	kospi200DayOverDay string
}

type KOSDAQ struct {
	kosdaq           string
	kosdaqUpDown     string
	kosdaqDayOverDay string
}

type KOSDAQStarIndex struct {
	kosdaqStarIndex           string
	kosdaqStarIndexUpDown     string
	kosdaqStarIndexDayOverDay string
}

type KRX100 struct {
	krx100           string
	krx100UpDown     string
	krx100DayOverDay string
}

func GetKOSPI() (*KOSPI, error) {
	s, err := getStockInfoBySymbol("000000")
	if err != nil {
		return nil, err
	}

	return &KOSPI{
		kospi:           s.StockInfo.KospiJisu,
		kospiUpDown:     s.StockInfo.KospiBuho,
		kospiDayOverDay: s.StockInfo.KospiDebi,
	}, nil
}

func GetKOSPI200() (*KOSPI200, error) {
	s, err := getStockInfoBySymbol("000000")
	if err != nil {
		return nil, err
	}

	return &KOSPI200{
		kospi200:           s.StockInfo.Kospi200Jisu,
		kospi200UpDown:     s.StockInfo.Kospi200Buho,
		kospi200DayOverDay: s.StockInfo.Kospi200Debi,
	}, nil
}

func GetKOSDAQ() (*KOSDAQ, error) {
	s, err := getStockInfoBySymbol("000000")
	if err != nil {
		return nil, err
	}

	return &KOSDAQ{
		kosdaq:           s.StockInfo.KosdaqJisu,
		kosdaqUpDown:     s.StockInfo.KosdaqJisuBuho,
		kosdaqDayOverDay: s.StockInfo.KosdaqJisuDebi,
	}, nil
}

func GetKOSDAQStarIndex() (*KOSDAQStarIndex, error) {
	s, err := getStockInfoBySymbol("000000")
	if err != nil {
		return nil, err
	}

	return &KOSDAQStarIndex{
		kosdaqStarIndex:           s.StockInfo.StarJisu,
		kosdaqStarIndexUpDown:     s.StockInfo.StarJisuBuho,
		kosdaqStarIndexDayOverDay: s.StockInfo.StarJisuDebi,
	}, nil
}

func GetKRX100() (*KRX100, error) {
	s, err := getStockInfoBySymbol("000000")
	if err != nil {
		return nil, err
	}

	return &KRX100{
		krx100:           s.StockInfo.Krx100Jisu,
		krx100UpDown:     s.StockInfo.Krx100buho,
		krx100DayOverDay: s.StockInfo.Krx100Debi,
	}, nil
}
