package amazon

import "time"

type ListOrdersResponse struct {
	ListOrdersResult struct {
	Orders struct {
	Order []struct {
	LatestShipDate         time.Time `json:"LatestShipDate"`
	OrderType              string    `json:"OrderType"`
	PurchaseDate           time.Time `json:"PurchaseDate"`
	PaymentExecutionDetail struct {
	PaymentExecutionDetailItem struct {
	PaymentMethod string `json:"PaymentMethod"`
	Payment       struct {
	CurrencyCode string `json:"CurrencyCode"`
	Amount       string `json:"Amount"`
} `json:"Payment"`
} `json:"PaymentExecutionDetailItem"`
} `json:"PaymentExecutionDetail,omitempty"`
	AmazonOrderID          string    `json:"AmazonOrderId"`
	BuyerEmail             string    `json:"BuyerEmail,omitempty"`
	IsReplacementOrder     string    `json:"IsReplacementOrder"`
	LastUpdateDate         time.Time `json:"LastUpdateDate"`
	NumberOfItemsShipped   string    `json:"NumberOfItemsShipped"`
	ShipServiceLevel       string    `json:"ShipServiceLevel"`
	OrderStatus            string    `json:"OrderStatus"`
	SalesChannel           string    `json:"SalesChannel"`
	ShippedByAmazonTFM     string    `json:"ShippedByAmazonTFM,omitempty"`
	IsBusinessOrder        string    `json:"IsBusinessOrder"`
	LatestDeliveryDate     time.Time `json:"LatestDeliveryDate,omitempty"`
	NumberOfItemsUnshipped string    `json:"NumberOfItemsUnshipped"`
	PaymentMethodDetails   struct {
	PaymentMethodDetail string `json:"PaymentMethodDetail"`
} `json:"PaymentMethodDetails"`
	BuyerName            string    `json:"BuyerName,omitempty"`
	EarliestDeliveryDate time.Time `json:"EarliestDeliveryDate,omitempty"`
	OrderTotal           struct {
	CurrencyCode string `json:"CurrencyCode"`
	Amount       string `json:"Amount"`
} `json:"OrderTotal,omitempty"`
	IsPremiumOrder     string    `json:"IsPremiumOrder"`
	EarliestShipDate   time.Time `json:"EarliestShipDate"`
	MarketplaceID      string    `json:"MarketplaceId"`
	FulfillmentChannel string    `json:"FulfillmentChannel"`
	TFMShipmentStatus  string    `json:"TFMShipmentStatus,omitempty"`
	PaymentMethod      string    `json:"PaymentMethod,omitempty"`
	ShippingAddress    struct {
	City          string `json:"City"`
	PostalCode    string `json:"PostalCode"`
	StateOrRegion string `json:"StateOrRegion"`
	CountryCode   string `json:"CountryCode"`
	Name          string `json:"Name"`
	AddressLine1  string `json:"AddressLine1"`
	AddressLine2  string `json:"AddressLine2"`
} `json:"ShippingAddress,omitempty"`
	IsPrime                      string `json:"IsPrime,omitempty"`
	ShipmentServiceLevelCategory string `json:"ShipmentServiceLevelCategory,omitempty"`
} `json:"Order"`
} `json:"Orders"`
} `json:"ListOrdersResult"`
}


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
