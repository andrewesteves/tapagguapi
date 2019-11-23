package handler

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/andrewesteves/tapagguapi/common"
	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/service"
	"github.com/gorilla/mux"
)

type ReceiptHttpHandler struct {
	Rs service.ReceiptContractService
}

func NewReceiptHttpHandler(mux *mux.Router, receiptService service.ReceiptContractService) {
	handler := &ReceiptHttpHandler{
		Rs: receiptService,
	}
	mux.HandleFunc("/receipts/retrieve", handler.Retrieve())
	mux.HandleFunc("/receipts/{id}", handler.Update()).Methods("PUT")
	mux.HandleFunc("/receipts/{id}", handler.Destroy()).Methods("DELETE")
	mux.HandleFunc("/receipts", handler.Store()).Methods("POST")
	mux.HandleFunc("/receipts", handler.All())
}

func (rh ReceiptHttpHandler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		receipts, err := rh.Rs.All()
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(receipts)
	}
}

func (rh ReceiptHttpHandler) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		receipt, err := rh.Rs.Find(int64(id))
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(receipt)
	}
}

func (rh ReceiptHttpHandler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var receipt model.Receipt
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &receipt)
		receipt, err := rh.Rs.Store(receipt)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(receipt)
	}
}

func (rh ReceiptHttpHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var receipt model.Receipt
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &receipt)
		receipt.ID = int64(id)
		receipt, err = rh.Rs.Update(receipt)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(receipt)
	}
}

func (rh ReceiptHttpHandler) Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		receipt, err := rh.Rs.Destroy(int64(id))
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(receipt)
	}
}

func (rh ReceiptHttpHandler) Retrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var receipt model.Receipt
		data, err := common.GetURL(r.URL.Query().Get("url"))
		if err != nil {
			log.Printf("Failed to get XML: %v", err)
		}
		xml.Unmarshal(data, &receipt)
		rpt, err := rh.Rs.Store(receipt)
		if err != nil {
			log.Printf("Failed to store: %v", err)
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rpt)
	}
}
