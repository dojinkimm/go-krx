package main

import (
	"sort"
)

type DailyStock struct {
	date             string
	endPrice         string
	upDown           string
	comparedToPrev   string
	marketPrice      string
	lowPrice         string
	highPrice        string
	volume           string
	transactionPrice string
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
			date:             d.DayDate,
			endPrice:         d.DayEndPrice,
			upDown:           d.DayDungrak,
			comparedToPrev:   d.DayGetDebi,
			marketPrice:      d.DayStart,
			lowPrice:         d.DayLow,
			highPrice:        d.DayHigh,
			volume:           d.DayVolume,
			transactionPrice: d.DayGetAmount,
		}
	}

	if len(dailyStock) < date {
		return dailyStock, nil
	}

	return dailyStock[:date], nil
}
