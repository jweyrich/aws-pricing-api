package schema

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type AWSConfig struct {
	FormatVersion	string
	Disclaimer	string
	OfferCode	string
	Version		string
	PublicationDate	string
	Products	map[string]AWSConfig_Product
	Terms		map[string]map[string]AWSConfig_Term
}
type AWSConfig_Product struct {	Sku	string
	ProductFamily	string
	Attributes	AWSConfig_Product_Attributes
}
type AWSConfig_Product_Attributes struct {	Operation	string
	Servicecode	string
	Location	string
	LocationType	string
	Usagetype	string
}

type AWSConfig_Term struct {
	OfferTermCode string
	Sku	string
	EffectiveDate string
	PriceDimensions AWSConfig_Term_PriceDimensions
	TermAttributes AWSConfig_Term_TermAttributes
}

type AWSConfig_Term_PriceDimensions struct {
	RateCode	string
	RateType	string
	Description	string
	BeginRange	string
	EndRange	string
	Unit	string
	PricePerUnit	AWSConfig_Term_PricePerUnit
	AppliesTo	[]interface{}
}

type AWSConfig_Term_PricePerUnit struct {
	USD	string
}

type AWSConfig_Term_TermAttributes struct {

}
func (a *AWSConfig) Refresh() error {
	var url = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AWSConfig/current/index.json"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, a)
	if err != nil {
		return err
	}

	return nil
}