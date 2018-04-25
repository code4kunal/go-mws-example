package amazon

import "time"

type ListOrderResponse struct {
	LatestShipDate               time.Time
	OrderType                    string
	PurchaseDate                 time.Time
	PaymentExecutionDetail       PaymentExecutionDetail
	AmazonOrderID                string
	BuyerEmail                   string
	IsReplacementOrder           string
	LastUpdateDate               time.Time
	NumberOfItemsShipped         string
	ShipServiceLevel             string
	OrderStatus                  string
	SalesChannel                 string
	ShippedByAmazonTFM           string
	IsBusinessOrder              string
	LatestDeliveryDate           time.Time
	NumberOfItemsUnshipped       string
	PaymentMethodDetails         PaymentMethodDetails
	BuyerName                    string
	EarliestDeliveryDate         time.Time
	OrderTotal                   OrderTotal
	IsPremiumOrder               string
	EarliestShipDate             time.Time
	MarketplaceID                string
	FulfillmentChannel           string
	TFMShipmentStatus            string
	PaymentMethod                string
	ShippingAddress              ShippingAddress
	IsPrime                      string
	ShipmentServiceLevelCategory string
}

type PaymentExecutionDetail struct {
	PaymentExecutionDetailItem PaymentExecutionDetailItem
}

type PaymentExecutionDetailItem struct {
	PaymentMethod string
	Payment       Payment
}

type Payment struct {
	CurrencyCode string
	Amount       string
}

type PaymentMethodDetails struct {
	PaymentMethodDetail string
}

type OrderTotal struct {
	CurrencyCode string
	Amount       string
}

type ShippingAddress struct {
	City          string
	PostalCode    string
	StateOrRegion string
	CountryCode   string
	Name          string
	AddressLine1  string
	AddressLine2  string
}
