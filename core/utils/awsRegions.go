package utils

import "strings"
import (

	"go-jwt-example/core/models/amazon"
)

func RegionByCountry(country string)  amazon.Region{
	for _, region := range Regions {
		if strings.EqualFold(region.Country, country) {
			return region
		}
	}
	panic("Invalid region, check your data")
}

var Regions = []amazon.Region{
	{"NA", "CA", "https://mws.amazonservices.ca/", "A2EUQ1WTGCTBG2"},
	{"NA", "US", "https://mws.amazonservices.com/", "ATVPDKIKX0DER"},
	{"EU", "DE", "https://mws-eu.amazonservices.com/", "A1PA6795UKMFR9"},
	{"EU", "ES", "https://mws-eu.amazonservices.com/", "A1RKKUPIHCS9HS"},
	{"EU", "FR", "https://mws-eu.amazonservices.com/", "A13V1IB3VIYZZH"},
	{"EU", "IN", "https://mws.amazonservices.in/", "A21TJRUUN4KGV"},
	{"EU", "IT", "https://mws-eu.amazonservices.com/", "APJ6JRA9NG5V4"},
	{"EU", "UK", "https://mws-eu.amazonservices.com/", "A1F83G8C2ARO7P"},
	{"FE", "JP", "https://mws.amazonservices.jp/", "A1VC38T7YXB528"},
	{"CN", "CN", "https://mws.amazonservices.com.cn/", "AAHKV2X7AFYLW"},
}
