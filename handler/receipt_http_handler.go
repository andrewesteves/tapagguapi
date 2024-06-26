package handler

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/andrewesteves/tapagguapi/common"
	"github.com/andrewesteves/tapagguapi/config"
	"github.com/andrewesteves/tapagguapi/middleware"
	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/service"
	"github.com/andrewesteves/tapagguapi/transformer"
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
		var wg sync.WaitGroup
		var receipts []model.Receipt
		var categories []model.Category
		var err error
		values := make(map[string]string)
		if r.URL.Query().Get("month") != "" {
			values["month"] = r.URL.Query().Get("month")
		}
		if r.URL.Query().Get("year") != "" {
			values["year"] = r.URL.Query().Get("year")
		}
		if r.URL.Query().Get("category") != "" {
			values["category"] = r.URL.Query().Get("category")
		}
		wg.Add(2)
		go func() {
			receipts, err = rh.Rs.All(*middleware.GetUser(r.Context()), values)
			if err != nil {
				panic(err.Error())
			}
			wg.Done()
		}()
		go func() {
			categories, err = rh.Rs.GroupCategoryTotal(*middleware.GetUser(r.Context()), values)
			if err != nil {
				panic(err.Error())
			}
			wg.Done()
		}()
		wg.Wait()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformer.ReceiptTransformer{}.TransformMany(receipts, categories, nil))
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

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformer.ReceiptTransformer{}.TransformOne(receipt))
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

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformer.ReceiptTransformer{}.TransformOne(receipt))
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
		receipt.User = *middleware.GetUser(r.Context())
		receipt, err = rh.Rs.Update(receipt)
		if err != nil {
			panic(err.Error())
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformer.ReceiptTransformer{}.TransformOne(receipt))
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

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformer.ReceiptTransformer{}.TransformOne(receipt))
	}
}

// Retrieve handler of a receipt
func (rh ReceiptHTTPHandler) Retrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var receipt model.Receipt
		data, err := common.GetURL(r.URL.Query().Get("url"))
		if err != nil {
			log.Printf("Failed to get XML: %v", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{
				"message": config.LangConfig{}.I18n()["process_failed"],
			})
			return
		}
		xml.Unmarshal(data, &receipt)
		if receipt.Total <= 0 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{
				"message": config.LangConfig{}.I18n()["receipt_notfound"],
			})
			return
		}
		receipt.User = *middleware.GetUser(r.Context())
		rpt, err := rh.Rs.RetrieveStore(receipt)
		if err != nil {
			log.Printf("Failed to store: %v", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{
				"message": config.LangConfig{}.I18n()["store_failed"],
			})
			return
		}
		dReceipt, err := rh.Rs.Find(rpt.ID)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformer.ReceiptTransformer{}.TransformOne(dReceipt))
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

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformer.ReceiptTransformer{}.TransformMany(receipts, nil, nil))
	}
}
