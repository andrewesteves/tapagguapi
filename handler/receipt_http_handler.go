package handler

import (
	"encoding/json"
	"net/http"

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
	mux.HandleFunc("/receipts", handler.All())
}

func (rh ReceiptHttpHandler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.Receipt{})
	}
}
