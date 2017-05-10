package schema

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type AmazonElastiCache struct {
	FormatVersion	string
	Disclaimer	string
	OfferCode	string
	Version		string
	PublicationDate	string
	Products	map[string]AmazonElastiCache_Product
	Terms		map[string]map[string]AmazonElastiCache_Term
}
type AmazonElastiCache_Product struct {	Sku	string
	ProductFamily	string
	Attributes	AmazonElastiCache_Product_Attributes
}
type AmazonElastiCache_Product_Attributes struct {	Servicecode	string
	Location	string
	Vcpu	string
	Memory	string
	NetworkPerformance	string
	Usagetype	string
	Operation	string
	LocationType	string
	InstanceType	string
	CurrentGeneration	string
	InstanceFamily	string
	CacheEngine	string
}

type AmazonElastiCache_Term struct {
	OfferTermCode string
	Sku	string
	EffectiveDate string
	PriceDimensions AmazonElastiCache_Term_PriceDimensions
	TermAttributes AmazonElastiCache_Term_TermAttributes
}

type AmazonElastiCache_Term_PriceDimensions struct {
	RateCode	string
	RateType	string
	Description	string
	BeginRange	string
	EndRange	string
	Unit	string
	PricePerUnit	AmazonElastiCache_Term_PricePerUnit
	AppliesTo	[]interface{}
}

type AmazonElastiCache_Term_PricePerUnit struct {
	USD	string
}

type AmazonElastiCache_Term_TermAttributes struct {

}
func (a *AmazonElastiCache) Refresh() error {
	var url = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AmazonElastiCache/current/index.json"
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