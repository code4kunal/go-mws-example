package api

import (
	"github.com/gorilla/schema"
	"github.com/gorilla/mux"
	"net/http"
	"go-jwt-example/core"
)


const (
	apiVersion = "1"
)

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