package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/service"
	"github.com/gorilla/mux"
)

type ItemHttpHandler struct {
	Is service.ItemContractService
}

func NewItemHttpHandler(mux *mux.Router, itemService service.ItemContractService) {
	handler := &ItemHttpHandler{
		Is: itemService,
	}
	mux.HandleFunc("/receipts/{receiptId}/items/{id}", handler.Find()).Methods("GET")
	mux.HandleFunc("/receipts/{receiptId}/items/{id}", handler.Update()).Methods("PUT")
	mux.HandleFunc("/receipts/{receiptId}/items/{id}", handler.Destroy()).Methods("DELETE")
	mux.HandleFunc("/receipts/{receiptId}/items", handler.Store()).Methods("POST")
	mux.HandleFunc("/receipts/{receiptId}/items", handler.All()).Methods("GET")
}

func (ih ItemHttpHandler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		receiptId, _ := strconv.Atoi(vars["receiptId"])
		items, err := ih.Is.All(int64(receiptId))
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(items)
	}
}

func (ih ItemHttpHandler) Find() http.HandlerFunc {
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

func (ih ItemHttpHandler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item model.Item
		vars := mux.Vars(r)
		receiptId, _ := strconv.Atoi(vars["receiptId"])
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &item)
		item, err := ih.Is.Store(int64(receiptId), item)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(item)
	}
}

func (ih ItemHttpHandler) Update() http.HandlerFunc {
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

func (ih ItemHttpHandler) Destroy() http.HandlerFunc {
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
