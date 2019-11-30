package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/service"
	"github.com/gorilla/mux"
)

// ItemHTTPHandler struct
type ItemHTTPHandler struct {
	Is service.ItemContractService
}

// NewItemHTTPHandler new item handler
func NewItemHTTPHandler(mux *mux.Router, itemService service.ItemContractService) {
	handler := &ItemHTTPHandler{
		Is: itemService,
	}
	mux.HandleFunc("/receipts/{receiptID}/items/{id}", handler.Find()).Methods("GET")
	mux.HandleFunc("/receipts/{receiptID}/items/{id}", handler.Update()).Methods("PUT")
	mux.HandleFunc("/receipts/{receiptID}/items/{id}", handler.Destroy()).Methods("DELETE")
	mux.HandleFunc("/receipts/{receiptID}/items", handler.Store()).Methods("POST")
	mux.HandleFunc("/receipts/{receiptID}/items", handler.All()).Methods("GET")
}

// All item handler
func (ih ItemHTTPHandler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Context().Value("token"))
		vars := mux.Vars(r)
		receiptID, _ := strconv.Atoi(vars["receiptID"])
		items, err := ih.Is.All(int64(receiptID))
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(items)
	}
}

// Find item handler
func (ih ItemHTTPHandler) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		item, err := ih.Is.Find(int64(id))
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(item)
	}
}

// Store item handler
func (ih ItemHTTPHandler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item model.Item
		vars := mux.Vars(r)
		receiptID, _ := strconv.Atoi(vars["receiptID"])
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &item)
		item, err := ih.Is.Store(int64(receiptID), item)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(item)
	}
}

// Update item handler
func (ih ItemHTTPHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item model.Item
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &item)
		item.ID = int64(id)
		item, err = ih.Is.Update(item)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(item)
	}
}

// Destroy item handler
func (ih ItemHTTPHandler) Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		item, err := ih.Is.Destroy(int64(id))
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(item)
	}
}
