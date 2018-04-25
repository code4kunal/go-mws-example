package services

import "net/url"
import (
	"go-jwt-example/core/utils"
	"time"
	"net/http"
	"errors"
	"strings"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"go-jwt-example/core/models/amazon"
)
const (

	LIST_ORDER_URL = "https://mws.amazonservices.in/Orders/2013-09-01"

)
var IncompleteRequest = errors.New("incomplete request")
type Client struct {
	AwsCreds         AwsCreds
	Method           string // GET, PUT, POST, etc.
	Region           amazon.Region
	Action           string
	Parameters       url.Values
	SignatureMethod  string
	SignatureVersion string
	CompanyName      string
}

type AwsCreds struct{
	AccessId  string
	AccessKey string
	MerchantId  string
	MarketPlaceId string
	MWSAuthToken string
}

func NewClient(client Client) Client {
	client.SignatureVersion = "2"
	client.SignatureMethod = "HmacSHA256"
	client.Region = utils.RegionByCountry("IN")
	if client.Parameters == nil {
		client.Parameters = make(url.Values)
	}
	return client
}

func (this *Client) Request() (req *http.Request, err error) {
	if this.Method == "" || this.AwsCreds.AccessId == "" || this.AwsCreds.AccessKey == "" ||
		this.AwsCreds.MerchantId == "" {
		err = IncompleteRequest
		return
	}

	this.Parameters.Add("Merchant", this.AwsCreds.MerchantId)
	this.Parameters.Add("AWSAccessKeyId", this.AwsCreds.AccessId)
	this.Parameters.Add("SignatureMethod", this.SignatureMethod)
	this.Parameters.Add("SignatureVersion", this.SignatureVersion)
	this.Parameters.Add("Version", "2009-01-01")
	this.Parameters.Add("Action", this.Action)
	this.Parameters.Add("Timestamp", XMLTimestamp(time.Now()))
	this.Region.Endpoint =  LIST_ORDER_URL
	stringToSign, err := this.StringToSign()
	if err != nil {
		return
	}

	url, err := url.Parse(this.Region.Endpoint)
	if err != nil {
		return
	}
	signature := Sign(stringToSign, []byte(this.AwsCreds.AccessKey))
	this.Parameters.Add("Signature", signature)
	url.RawQuery = CanonicalizedQueryString(this.Parameters)
	req, err = http.NewRequest(this.Method, url.String(), nil)
	//req.Header.Add("User-Agent", UserAgent)
	return
}

var ISO8601 = "2006-01-02T15:04:05Z"

func XMLTimestamp(t time.Time) string {
	return t.UTC().Format(ISO8601)
}

func CanonicalizedQueryString(values url.Values) (str string) {
	// per aws docs and docs for values.Encode, we respect RFC 3986
	// we may not deal with utf-8, only ascii
	// params are sorted
	// we have to fix the '+' to '%20'
	str = values.Encode()
	str = strings.Replace(str, "+", "%20", -1)
	return
}
func (this *Client) StringToSign() (stringToSign string, err error) {
	endpoint, err := url.Parse(this.Region.Endpoint)
	if err != nil {
		return
	}
	stringToSign = strings.Join([]string{
		this.Method,
		strings.ToLower(endpoint.Host),
		endpoint.Path,
		CanonicalizedQueryString(this.Parameters),
	}, "\n")

	return
}

func Sign(str string, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

//func GetOrders(r *http.Request)( error){
//	resp, err := resty.R().
//		SetQueryParams(map[string]string{
//		"Action": "ListOrders",
//		"SellerId": client.AwsCreds.MerchantId,
//		"MWSAuthToken":client.AwsCreds.MWSAuthToken,
//		"SignatureVersion": "2",
//		"Timestamp":strconv.FormatInt(time.Now().Unix(), 10),
//		"Version":"2013-09-01",
//		""
//	}).
//		SetHeader("Accept", "application/json").
//		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
//		Get("/show_product")
//
//}