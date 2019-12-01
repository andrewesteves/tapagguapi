package handler

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/andrewesteves/tapagguapi/common"
	"github.com/andrewesteves/tapagguapi/middleware"
	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/service"
	"github.com/andrewesteves/tapagguapi/transformers"
	"github.com/gorilla/mux"
)

// ReceiptHTTPHandler struct
type ReceiptHTTPHandler struct {
	Rs service.ReceiptContractService
}

// NewReceiptHTTPHandler new receipt handler
func NewReceiptHTTPHandler(mux *mux.Router, receiptService service.ReceiptContractService) {
	handler := &ReceiptHTTPHandler{
		Rs: receiptService,
	}
	mux.HandleFunc("/receipts/query/{field}", handler.Query()).Methods("GET")
	mux.HandleFunc("/receipts/retrieve", handler.Retrieve()).Methods("GET")
	mux.HandleFunc("/receipts/{id}", handler.Find()).Methods("GET")
	mux.HandleFunc("/receipts/{id}", handler.Update()).Methods("PUT")
	mux.HandleFunc("/receipts/{id}", handler.Destroy()).Methods("DELETE")
	mux.HandleFunc("/receipts", handler.Store()).Methods("POST")
	mux.HandleFunc("/receipts", handler.All()).Methods("GET")
}

// All handler of a receipts
func (rh ReceiptHTTPHandler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		receipts, err := rh.Rs.All(*middleware.GetUser(r.Context()))
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformers.ReceiptTransformer{}.TransformMany(receipts, nil))
	}
}

// Find handler of a receipt
func (rh ReceiptHTTPHandler) Find() http.HandlerFunc {
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

// Store handler of a receipt
func (rh ReceiptHTTPHandler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var receipt model.Receipt
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &receipt)
		receipt.User = *middleware.GetUser(r.Context())
		receipt, err := rh.Rs.Store(receipt)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(receipt)
	}
}

// Update handler of a receipt
func (rh ReceiptHTTPHandler) Update() http.HandlerFunc {
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

// Destroy handler of a receipt
func (rh ReceiptHTTPHandler) Destroy() http.HandlerFunc {
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

// Retrieve handler of a receipt
func (rh ReceiptHTTPHandler) Retrieve() http.HandlerFunc {
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

// Query handler of a receipt
func (rh ReceiptHTTPHandler) Query() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		receipts, err := rh.Rs.FindManyBy(vars["field"], r.URL.Query().Get("value"))
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformers.ReceiptTransformer{}.TransformMany(receipts, nil))
	}
}
