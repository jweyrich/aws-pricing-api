package schema

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type AmazonGameLift struct {
	FormatVersion	string
	Disclaimer	string
	OfferCode	string
	Version		string
	PublicationDate	string
	Products	map[string]AmazonGameLift_Product
	Terms		map[string]map[string]AmazonGameLift_Term
}
type AmazonGameLift_Product struct {	Sku	string
	ProductFamily	string
	Attributes	AmazonGameLift_Product_Attributes
}
type AmazonGameLift_Product_Attributes struct {	Servicecode	string
	TransferType	string
	FromLocation	string
	FromLocationType	string
	ToLocation	string
	ToLocationType	string
	Usagetype	string
	Operation	string
}

type AmazonGameLift_Term struct {
	OfferTermCode string
	Sku	string
	EffectiveDate string
	PriceDimensions AmazonGameLift_Term_PriceDimensions
	TermAttributes AmazonGameLift_Term_TermAttributes
}

type AmazonGameLift_Term_PriceDimensions struct {
	RateCode	string
	RateType	string
	Description	string
	BeginRange	string
	EndRange	string
	Unit	string
	PricePerUnit	AmazonGameLift_Term_PricePerUnit
	AppliesTo	[]interface{}
}

type AmazonGameLift_Term_PricePerUnit struct {
	USD	string
}

type AmazonGameLift_Term_TermAttributes struct {

}
func (a *AmazonGameLift) Refresh() error {
	var url = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AmazonGameLift/current/index.json"
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