package api

import (
	"github.com/gorilla/mux"

)

func wireupRoutes(r *mux.Router, api *API) {

	r.HandleFunc("/api", api.fetchInfo).Methods("GET")
	r.HandleFunc("/api/v1/login", api.login).Methods("POST")
	r.HandleFunc("/api/v1/logout", api.logout).Methods("GET")
	r.HandleFunc("/api/v1/sync/orders", api.parseRequestAndCreateInvoice).Methods("POST")
	//r.HandleFunc("/api/v1/reports/activities", api.activitiesOverview).Methods("GET")

	//return
}

