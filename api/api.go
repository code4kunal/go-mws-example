package api

import (
	"github.com/gorilla/schema"
	"github.com/gorilla/mux"
	"net/http"
	"go-jwt-example/core"
	"github.com/ezkl/go-amazon-mws-api"
	"go-jwt-example/core/services"
	"errors"
	"fmt"
	"io/ioutil"
)


const (
	apiVersion = "1"
)

var awsApi amazonmws.AmazonMWSAPI
// API Container object
type API struct {
	AppVersion string
	c          *core.Core
	reqDecoder *schema.Decoder
}

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
 	 AwsClient := services.NewClient(client)
	 creds:= services.AwsCreds{
		 AccessId: "AKIAJEHLC4BUI5SKV3PA",
		 AccessKey: "Z/Rs9NMMv4wBrlSQvPWBeEhszuGFYaAh596F/Crt",
		 MerchantId: "A17LG0A22TE4YC",
		 MarketPlaceId: "A21TJRUUN4KGV",
		 MWSAuthToken: "amzn.mws.bcb17b76-c55b-3243-86c4-535f72857242",
	 }
   AwsClient.AwsCreds = creds
   fmt.Println(AwsClient)
   req, err := AwsClient.Request()
   if(err!= nil){
   	renderError(w, errors.New("Bad Request"), 400)
   }

	awsPostClient := &http.Client{}
	resp, err := awsPostClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}