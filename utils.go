package krx

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StockInformation struct {
	XMLName       xml.Name `xml:"stockprice"`
	Text          string   `xml:",chardata"`
	Querytime     string   `xml:"querytime,attr"`
	TBLDailyStock struct {
		Text       string `xml:",chardata"`
		DailyStock []struct {
			Text         string `xml:",chardata"`
			DayDate      string `xml:"day_Date,attr"`
			DayEndPrice  string `xml:"day_EndPrice,attr"`
			DayDungrak   string `xml:"day_Dungrak,attr"`
			DayGetDebi   string `xml:"day_getDebi,attr"`
			DayStart     string `xml:"day_Start,attr"`
			DayHigh      string `xml:"day_High,attr"`
			DayLow       string `xml:"day_Low,attr"`
			DayVolume    string `xml:"day_Volume,attr"`
			DayGetAmount string `xml:"day_getAmount,attr"`
		}
	} `xml:"TBL_DailyStock"`
	TBLAskPrice struct {
		Text     string `xml:",chardata"`
		AskPrice []struct {
			Text           string `xml:",chardata"`
			MemberMemdoMem string `xml:"member_memdoMem,attr"`
			MemberMemdoVol string `xml:"member_memdoVol,attr"`
			MemberMemsoMem string `xml:"member_memsoMem,attr"`
			MemberMesuoVol string `xml:"member_mesuoVol,attr"`
		} `xml:"AskPrice"`
	} `xml:"TBL_AskPrice"`
	TBLStockInfo struct {
		Text      string `xml:",chardata"`
		JongName  string `xml:"JongName,attr"`
		CurJuka   string `xml:"CurJuka,attr"`
		DungRak   string `xml:"DungRak,attr"`
		Debi      string `xml:"Debi,attr"`
		PrevJuka  string `xml:"PrevJuka,attr"`
		Volume    string `xml:"Volume,attr"`
		Money     string `xml:"Money,attr"`
		StartJuka string `xml:"StartJuka,attr"`
		HighJuka  string `xml:"HighJuka,attr"`
		LowJuka   string `xml:"LowJuka,attr"`
		High52    string `xml:"High52,attr"`
		Low52     string `xml:"Low52,attr"`
		UpJuka    string `xml:"UpJuka,attr"`
		DownJuka  string `xml:"DownJuka,attr"`
		Per       string `xml:"Per,attr"`
		Amount    string `xml:"Amount,attr"`
		FaceJuka  string `xml:"FaceJuka,attr"`
	} `xml:"TBL_StockInfo"`
	TBLHoga struct {
		Text      string `xml:",chardata"`
		MesuJan0  string `xml:"mesuJan0,attr"`
		MesuHoka0 string `xml:"mesuHoka0,attr"`
		MesuJan1  string `xml:"mesuJan1,attr"`
		MesuHoka1 string `xml:"mesuHoka1,attr"`
		MesuJan2  string `xml:"mesuJan2,attr"`
		MesuHoka2 string `xml:"mesuHoka2,attr"`
		MesuJan3  string `xml:"mesuJan3,attr"`
		MesuHoka3 string `xml:"mesuHoka3,attr"`
		MesuJan4  string `xml:"mesuJan4,attr"`
		MesuHoka4 string `xml:"mesuHoka4,attr"`
		MedoJan0  string `xml:"medoJan0,attr"`
		MedoHoka0 string `xml:"medoHoka0,attr"`
		MedoJan1  string `xml:"medoJan1,attr"`
		MedoHoka1 string `xml:"medoHoka1,attr"`
		MedoJan2  string `xml:"medoJan2,attr"`
		MedoHoka2 string `xml:"medoHoka2,attr"`
		MedoJan3  string `xml:"medoJan3,attr"`
		MedoHoka3 string `xml:"medoHoka3,attr"`
		MedoJan4  string `xml:"medoJan4,attr"`
		MedoHoka4 string `xml:"medoHoka4,attr"`
	} `xml:"TBL_Hoga"`
	TBLTimeConclude struct {
		Text            string `xml:",chardata"`
		TBLTimeConclude []struct {
			Text      string `xml:",chardata"`
			Time      string `xml:"time,attr"`
			Negoprice string `xml:"negoprice,attr"`
			Dungrak   string `xml:"Dungrak,attr"`
			Debi      string `xml:"Debi,attr"`
			Sellprice string `xml:"sellprice,attr"`
			Buyprice  string `xml:"buyprice,attr"`
			Amount    string `xml:"amount,attr"`
		} `xml:"TBL_TimeConclude"`
	} `xml:"TBL_TimeConclude"`
	StockInfo struct {
		Text           string `xml:",chardata"`
		KosdaqJisu     string `xml:"kosdaqJisu,attr"`
		KosdaqJisuBuho string `xml:"kosdaqJisuBuho,attr"`
		KosdaqJisuDebi string `xml:"kosdaqJisuDebi,attr"`
		StarJisu       string `xml:"starJisu,attr"`
		StarJisuBuho   string `xml:"starJisuBuho,attr"`
		StarJisuDebi   string `xml:"starJisuDebi,attr"`
		Jisu50         string `xml:"jisu50,attr"`
		Jisu50Buho     string `xml:"jisu50Buho,attr"`
		Jisu50Debi     string `xml:"jisu50Debi,attr"`
		MyNowTime      string `xml:"myNowTime,attr"`
		MyJangGubun    string `xml:"myJangGubun,attr"`
		MyPublicPrice  string `xml:"myPublicPrice,attr"`
		Krx100Jisu     string `xml:"krx100Jisu,attr"`
		Krx100buho     string `xml:"krx100buho,attr"`
		Krx100Debi     string `xml:"krx100Debi,attr"`
		KospiJisu      string `xml:"kospiJisu,attr"`
		KospiBuho      string `xml:"kospiBuho,attr"`
		KospiDebi      string `xml:"kospiDebi,attr"`
		Kospi200Jisu   string `xml:"kospi200Jisu,attr"`
		Kospi200Buho   string `xml:"kospi200Buho,attr"`
		Kospi200Debi   string `xml:"kospi200Debi,attr"`
	} `xml:"stockInfo"`
}

type Company struct {
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Sector          string `json:"sector"`
	Industry        string `json:"industry"`
	ListingDate     string `json:"listing_date"`
	SettlementMonth string `json:"settlement_month"`
	Representative  string `json:"representative"`
	Homepage        string `json:"homepage"`
	Region          string `json:"region"`
}

type UpDown int

const (
	UpperLimit UpDown = iota
	Increase
	Keep
	LowerLimit
	Decrease
	Unknown
)

func getStockInfoBySymbol(symbol string) (*StockInformation, error) {
	url := "http://asp1.krx.co.kr/servlet/krx.asp.XMLSiseEng?code="
	resp, err := http.Get(url + symbol)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Error(err)
		}
	}()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var stockInfo *StockInformation
	xmlerr := xml.Unmarshal(bodyBytes, &stockInfo)
	if xmlerr != nil {
		return nil, errors.WithStack(xmlerr)
	}

	if len(stockInfo.TBLDailyStock.DailyStock) == 0 {
		return nil, status.Errorf(codes.NotFound, "Stock not found by given symbol: %s", symbol)
	}

	return stockInfo, nil
}

func getStockInfoByCompanyName(companyName string) (*StockInformation, error) {
	symbol, err := GetSymbolByCompanyName(companyName)
	if err != nil {
		return nil, err
	}

	stockInfo, err := getStockInfoBySymbol(symbol)
	if err != nil {
		return nil, err
	}

	return stockInfo, nil
}

func GetSymbolByCompanyName(name string) (string, error) {
	companyInfo, err := getCompanyListFromJsonFile()
	if err != nil {
		return "", err
	}

	// TODO @dojinkimm - find element more efficiently (e.g. binary search)
	for _, c := range companyInfo {
		if c.Name == name {
			return c.Symbol, nil
		}
	}

	return "", status.Errorf(codes.NotFound, "Symbol is not found by given company name: %s", name)
}

func getCompanyListFromJsonFile() ([]*Company, error) {
	jsonFile, err := os.Open("data/data.json")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			log.Error(err)
		}
	}()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var companies []*Company
	if err = json.Unmarshal(byteValue, &companies); err != nil {
		return nil, errors.WithStack(err)
	}

	return companies, nil
}

func castToUpDown(num string) UpDown {
	switch num {
	case "1":
		return UpperLimit
	case "2":
		return Increase
	case "3":
		return Keep
	case "4":
		return LowerLimit
	case "5":
		return Decrease
	}

	return Unknown
}
