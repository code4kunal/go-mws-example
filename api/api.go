package api

import (
	"github.com/gorilla/schema"
	"github.com/gorilla/mux"
	"net/http"
	"go-jwt-example/core"
	"go-jwt-example/core/services"
	"errors"
	"fmt"

	"io/ioutil"
	"bytes"
	"encoding/xml"
	"gopkg.in/xmlpath.v2"
)

const (
	apiVersion                 = "1"
	LIST_ORDER                 = "ListOrders"
	REQUEST_REPORT             = "RequestReport"
	GET_REPORT_LIST            = "GetReportList"
	GET_REPORT                 = "GetReport"
	POST                       = "POST"
	GET                        = "GET"
	CREATE_INVOICE_SOCKET_FLOW = "https://sokt.io/FH4S7ryuTn7u5SLwbVgu/amazon-flow-amazon-create-invoice"
)

// API Container object
type API struct {
	AppVersion string
	c          *core.Core
	reqDecoder *schema.Decoder
}

type Envelope interface{}

// New instantiates our router that can then be used by an external request
// handler
func New(c *core.Core, r *mux.Router, version string) *API {
	a := &API{}
	a.AppVersion = version
	a.c = c
	a.reqDecoder = schema.NewDecoder()
	wireupRoutes(r, a)
	return a
}

// fetchInfo returns app configuration information
func (api *API) fetchInfo(w http.ResponseWriter, r *http.Request) {
	data := struct {
		AppVersion string `json:"app_version"`
	}{
		AppVersion: api.AppVersion,
	}
	renderJSON(w, data, http.StatusOK)
}

func (api *API) login(w http.ResponseWriter, r *http.Request) {

}

func (api *API) logout(w http.ResponseWriter, r *http.Request) {

}

func (api *API) validate(w http.ResponseWriter, r *http.Request) {

}

func (api *API) parseRequestAndCreateInvoice(w http.ResponseWriter, r *http.Request) {
	var client services.Client
	client.Operation = LIST_ORDER
	client.Action = LIST_ORDER
	client.Method = POST
	AwsClient := services.NewClient(client)
	creds := services.AwsCreds{
		AccessId:      "AKIAJEHLC4BUI5SKV3PA",
		AccessKey:     "Z/Rs9NMMv4wBrlSQvPWBeEhszuGFYaAh596F/Crt",
		MerchantId:    "A17LG0A22TE4YC",
		MarketPlaceId: "A21TJRUUN4KGV",
		MWSAuthToken:  "amzn.mws.bcb17b76-c55b-3243-86c4-535f72857242",
	}
	AwsClient.AwsCreds = creds

	req, err := AwsClient.Request()
	if (err != nil) {
		fmt.Println(err)
		renderError(w, errors.New("Bad Request"), 400)
	}

	awsPostClient := &http.Client{}
	clientForScoket := &http.Client{}
	resp, err := awsPostClient.Do(req)
	//url, err := url.Parse(CREATE_INVOICE_SOCKET_FLOW)
	var f interface{}
	body, _ := ioutil.ReadAll(resp.Body)
	//var s = new(amazon.ListOrdersResponse)
	err = xml.Unmarshal(body, &f)
	if err != nil {
		fmt.Println(err)
	}
	//in := bytes.NewReader(body)
	//req, er = sling.New().Base(CREATE_INVOICE_SOCKET_FLOW).Set("Content-Type", "text/plain").Body(in).Request()
	req, err = http.NewRequest("POST", CREATE_INVOICE_SOCKET_FLOW, bytes.NewReader(body))
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/xml") //req, err = http.NewRequest(POST, url.String(), body)
	clientForScoket.Do(req)
	defer resp.Body.Close()

}

func (api *API) parseRequestAndCreateStock(w http.ResponseWriter, r *http.Request) {

	var client services.Client
	client.Operation = REQUEST_REPORT
	client.Action = REQUEST_REPORT
	client.Method = POST
	AwsClient := services.NewClient(client)
	creds := services.AwsCreds{
		AccessId:      "AKIAJEHLC4BUI5SKV3PA",
		AccessKey:     "Z/Rs9NMMv4wBrlSQvPWBeEhszuGFYaAh596F/Crt",
		MerchantId:    "A17LG0A22TE4YC",
		MarketPlaceId: "A21TJRUUN4KGV",
		MWSAuthToken:  "amzn.mws.bcb17b76-c55b-3243-86c4-535f72857242",
	}
	AwsClient.AwsCreds = creds

	request, err := AwsClient.RequestForReport()
	if (err != nil) {
		fmt.Println(err)
		renderError(w, errors.New("Bad Request"), 400)
	}
	awsPostClient := &http.Client{}
	resp, err := awsPostClient.Do(request)
	body, _ := ioutil.ReadAll(resp.Body)
	path := xmlpath.MustCompile("/RequestReportResponse/RequestReportResult/ReportRequestInfo/ReportRequestId")
	root, err := xmlpath.Parse(bytes.NewReader(body))
	if err != nil {
		//
	}
	repId := ""
	if value, ok := path.String(root); ok {
		repId = value
	}
	reportId := repId
	finalReportId := services.CallListReportApi(reportId, creds)
	//if(err==nil){
	//	fmt.Println(err)
	//	fmt.Println("errror in request")
	//}



	////todo: here implement xpath for requestreportlist
	AwsClientForReport := services.NewClient(client)
	AwsClientForReport.AwsCreds = creds
	req, err := AwsClientForReport.GetReport(finalReportId)
	if (err != nil) {
		renderError(w, errors.New("Bad Request"), 400)
	}
	resp, err = awsPostClient.Do(req)
	body, _ = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)
	//
	//in := bytes.NewReader(body)
	////req, er = sling.New().Base(CREATE_INVOICE_SOCKET_FLOW).Set("Content-Type", "text/plain").Body(in).Request()
	//req, err = http.NewRequest("POST", CREATE_INVOICE_SOCKET_FLOW, bytes.NewReader(body))
	//req.Body = ioutil.NopCloser(bytes.NewReader(body))
	//req.Header.Add("Content-Type", "application/xml")	//req, err = http.NewRequest(POST, url.String(), body)
	//clientForScoket.Do(req)
	defer resp.Body.Close()

}
