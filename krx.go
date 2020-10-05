package krx

type KOSPI struct {
	KOSPIIndex      string
	KOSPIUpDown     string
	KOSPIDayOverDay string
}

type KOSPI200 struct {
	KOSPI200           string
	KOSPI200UpDown     string
	KOSPI200DayOverDay string
}

type KOSDAQ struct {
	KOSDAQIndex      string
	KOSDAQUpDown     string
	KOSDAQDayOverDay string
}

type KOSDAQStarIndex struct {
	KOSDAQStarIndex           string
	KOSDAQStarIndexUpDown     string
	KOSDAQStarIndexDayOverDay string
}

type KRX100 struct {
	KRX100Index      string
	KRX100UpDown     string
	KRX100DayOverDay string
}

func GetKOSPI() (*KOSPI, error) {
	s, err := getKRXMarketInfo()
	if err != nil {
		return nil, err
	}

	return &KOSPI{
		KOSPIIndex:      s.StockInfo.KospiJisu,
		KOSPIUpDown:     s.StockInfo.KospiBuho,
		KOSPIDayOverDay: s.StockInfo.KospiDebi,
	}, nil
}

func GetKOSPI200() (*KOSPI200, error) {
	s, err := getKRXMarketInfo()
	if err != nil {
		return nil, err
	}

	return &KOSPI200{
		KOSPI200:           s.StockInfo.Kospi200Jisu,
		KOSPI200UpDown:     s.StockInfo.Kospi200Buho,
		KOSPI200DayOverDay: s.StockInfo.Kospi200Debi,
	}, nil
}

func GetKOSDAQ() (*KOSDAQ, error) {
	s, err := getKRXMarketInfo()
	if err != nil {
		return nil, err
	}

	return &KOSDAQ{
		KOSDAQIndex:      s.StockInfo.KosdaqJisu,
		KOSDAQUpDown:     s.StockInfo.KosdaqJisuBuho,
		KOSDAQDayOverDay: s.StockInfo.KosdaqJisuDebi,
	}, nil
}

func GetKOSDAQStarIndex() (*KOSDAQStarIndex, error) {
	s, err := getKRXMarketInfo()
	if err != nil {
		return nil, err
	}

	return &KOSDAQStarIndex{
		KOSDAQStarIndex:           s.StockInfo.StarJisu,
		KOSDAQStarIndexUpDown:     s.StockInfo.StarJisuBuho,
		KOSDAQStarIndexDayOverDay: s.StockInfo.StarJisuDebi,
	}, nil
}

func GetKRX100() (*KRX100, error) {
	s, err := getKRXMarketInfo()
	if err != nil {
		return nil, err
	}

	return &KRX100{
		KRX100Index:      s.StockInfo.Krx100Jisu,
		KRX100UpDown:     s.StockInfo.Krx100buho,
		KRX100DayOverDay: s.StockInfo.Krx100Debi,
	}, nil
}
