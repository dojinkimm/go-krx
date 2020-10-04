package main

type Member struct {
	name   string
	volume string
}

func GetTopBuyersBySymbol(symbol string) ([]*Member, error) {
	s, err := getStockInfoBySymbol(symbol)
	if err != nil {
		return nil, err
	}

	members := make([]*Member, len(s.TBLAskPrice.AskPrice))
	for idx, d := range s.TBLAskPrice.AskPrice {
		members[idx] = &Member{
			name:   d.MemberMemsoMem,
			volume: d.MemberMesuoVol,
		}
	}

	return members, nil
}
func GetTopSellersBySymbol(symbol string) ([]*Member, error) {
	s, err := getStockInfoBySymbol(symbol)
	if err != nil {
		return nil, err
	}

	members := make([]*Member, len(s.TBLAskPrice.AskPrice))
	for idx, d := range s.TBLAskPrice.AskPrice {
		members[idx] = &Member{
			name:   d.MemberMemdoMem,
			volume: d.MemberMemdoVol,
		}
	}

	return members, nil
}
