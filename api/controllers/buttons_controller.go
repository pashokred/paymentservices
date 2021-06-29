package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"paymentservices/api/domain/button_domain"
	"paymentservices/api/services/base_service"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetButtons(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	request := button_domain.ButtonRequest{
		ProductID: vars["id"],
	}
	start := time.Now()
	result, apiError := base_service.BaseService.GetButtons(request)
	w.Header().Set("X-Response-Time", strconv.FormatInt(int64(time.Since(start)), 10))
	w.Header().Set("Content-Type", "application/json")

	hostname, err := os.Hostname()
	if err == nil {
		w.Header().Set("X-Server-Name", hostname)
	}

	if apiError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorJson, err := json.Marshal(apiError)
		if err != nil {
			log.Panic(err)
			return
		}
		_, err = w.Write(errorJson)
		if err != nil {
			log.Panic(err)
			return
		}
		log.Panic(err)
		return
	}
	buttonJson, err := json.Marshal(result)
	if err != nil {
		log.Panic(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(buttonJson)
	if err != nil {
		log.Panic(err)
		return
	}
}
