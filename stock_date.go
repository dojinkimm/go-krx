package main

import (
	"sort"
)

type DailyStock struct {
	date         string
	endPrice     string
	upDown       string
	dayOverDay   string
	marketValue  string
	low          string
	high         string
	volume       string
	tradingValue string
}

func GetStockByDateSymbol(symbol string, date int) ([]*DailyStock, error) {
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
			date:         d.DayDate,
			endPrice:     d.DayEndPrice,
			upDown:       d.DayDungrak,
			dayOverDay:   d.DayGetDebi,
			marketValue:  d.DayStart,
			low:          d.DayLow,
			high:         d.DayHigh,
			volume:       d.DayVolume,
			tradingValue: d.DayGetAmount,
		}
	}

	if len(dailyStock) < date {
		return dailyStock, nil
	}

	return dailyStock[:date], nil
}
