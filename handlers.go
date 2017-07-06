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
		log.Printf("%v", err)
	}

	log.Printf("%v", result)

}

func merchantSaveDummyData(w http.ResponseWriter, r *http.Request) {

	for i := 0; i < 10000; i++ {
		merchantID := strconv.Itoa(i)
		result, err := redis.NewRedis().Set(merchantID, getCurrentTimestamp())

		if err != nil {
			log.Printf("%v", err)
		}

		log.Printf("Result for key %s : %v", merchantID, result)
	}

}

func getMerchantLastUpdate(w http.ResponseWriter, r *http.Request) {
	merchantID := mux.Vars(r)["merchant_id"]
	result, err := redis.NewRedis().Get(merchantID)

	if err != nil {
		w.WriteHeader(404)
		log.Printf("%v", err)
	}

	if err2 := json.NewEncoder(w).Encode(result); err2 != nil {
		w.WriteHeader(500)
		log.Printf("%v", err)
	}

}

func getCurrentTimestamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}
