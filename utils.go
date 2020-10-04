package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type CompanyInformation struct {
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

func GetCodeNumberByCompanyName(name string) (string, error) {
	companyInfo, err := getCompanyInformationListFromJsonFile()
	if err != nil {
		return "", err
	}

	// TODO @dojinkimm - find element more efficiently (e.g. binary search)
	codeNumber := ""
	for _, c := range companyInfo {
		if c.Name == name {
			codeNumber = c.Symbol
			break
		}
	}

	return codeNumber, nil
}

func getCompanyInformationListFromJsonFile() ([]*CompanyInformation, error) {
	jsonFile, err := os.Open("data/data.json")
	if err != nil {
		log.Panic(err)
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

	var companyInfo []*CompanyInformation
	if err = json.Unmarshal(byteValue, &companyInfo); err != nil {
		return nil, errors.WithStack(err)
	}

	return companyInfo, nil
}
