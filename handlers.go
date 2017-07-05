package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"strconv"

	"github.com/arthurstockler/order-go-library/redis"
	"github.com/gorilla/mux"
)

func merchantSaveOrUpdate(w http.ResponseWriter, r *http.Request) {

	merchantID := r.Header.Get("merchant-id")
	result, err := redis.NewRedis().Set(merchantID, getCurrentTimestamp())

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", result)

}

func getMerchantLastUpdate(w http.ResponseWriter, r *http.Request) {
	merchantID := mux.Vars(r)["merchant_id"]
	result, err := redis.NewRedis().Get(merchantID)

	if err != nil {
		log.Fatal(err)
	}

	if err2 := json.NewEncoder(w).Encode(result); err2 != nil {
		panic(err2)
	}

}

func getCurrentTimestamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}
