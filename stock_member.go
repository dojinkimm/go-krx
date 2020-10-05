package krx

type Member struct {
	Name   string
	Volume string
}

func GetTopBuyersBySymbol(symbol string) ([]*Member, error) {
	s, err := getStockInfoBySymbol(symbol)
	if err != nil {
		return nil, err
	}

	members := make([]*Member, len(s.TBLAskPrice.AskPrice))
	for idx, d := range s.TBLAskPrice.AskPrice {
		members[idx] = &Member{
			Name:   d.MemberMemsoMem,
			Volume: d.MemberMesuoVol,
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
			Name:   d.MemberMemdoMem,
			Volume: d.MemberMemdoVol,
		}
	}

	return members, nil
}
