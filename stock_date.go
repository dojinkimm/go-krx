package krx

import (
	"sort"
)

type DailyStock struct {
	Date         string
	EndPrice     string
	UpDown       UpDown
	DayOverDay   string
	Volume       string
	TradingValue string
	StartValue   string
	Low          string
	High         string
}

func GetStockBySymbolAndDate(symbol string, date int) ([]*DailyStock, error) {
	s, err := getStockInfoBySymbol(symbol)
	if err != nil {
		return nil, err
	}

	sort.Slice(s.TBLDailyStock.DailyStock, func(i, j int) bool {
		return s.TBLDailyStock.DailyStock[i].DayDate > s.TBLDailyStock.DailyStock[j].DayDate
	})

	dailyStock := make([]*DailyStock, len(s.TBLDailyStock.DailyStock))
	for idx, d := range s.TBLDailyStock.DailyStock {
		dailyStock[idx] = &DailyStock{
			Date:         d.DayDate,
			EndPrice:     d.DayEndPrice,
			UpDown:       castToUpDown(d.DayDungrak),
			DayOverDay:   d.DayGetDebi,
			Volume:       d.DayVolume,
			TradingValue: d.DayGetAmount,
			StartValue:   d.DayStart,
			Low:          d.DayLow,
			High:         d.DayHigh,
		}
	}

	if len(dailyStock) < date {
		return dailyStock, nil
	}

	return dailyStock[:date], nil
}
