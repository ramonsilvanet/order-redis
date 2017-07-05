package routes

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/git-m4u/lio-order-manager-service-go/config"

	"github.com/gorilla/mux"
)

type Router struct {
	Rhttp *config.YmlHTTP
}

func (r *Router) ServerStart() {
	router := mux.NewRouter().StrictSlash(true)
	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%d", r.Rhttp.Host, r.Rhttp.Port),
	}
	fmt.Printf("Starting http server %v", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func apiV1(r *Router, router *mux.Router) {
	api := router.PathPrefix("/api/v1").Subrouter()
	apiMerchant(r, api)
}

func apiMerchant(r *Router, api *mux.Router) {

}
