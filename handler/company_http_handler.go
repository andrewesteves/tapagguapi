package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/andrewesteves/tapagguapi/middleware"
	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/service"
	"github.com/gorilla/mux"
)

// CompanyHTTPHandler struct
type CompanyHTTPHandler struct {
	Cs service.CompanyContractService
}

// NewCompanyHTTPHandler new company handler
func NewCompanyHTTPHandler(mux *mux.Router, companyService service.CompanyContractService) {
	handler := &CompanyHTTPHandler{
		Cs: companyService,
	}
	mux.HandleFunc("/companies/{id}", handler.Find()).Methods("GET")
	mux.HandleFunc("/companies/{id}", handler.Update()).Methods("PUT")
	mux.HandleFunc("/companies/{id}", handler.Destroy()).Methods("DELETE")
	mux.HandleFunc("/companies", handler.Store()).Methods("POST")
	mux.HandleFunc("/companies", handler.All()).Methods("GET")
}

// All handler of a companies
func (rh CompanyHTTPHandler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		companies, err := rh.Cs.All(*middleware.GetUser(r.Context()))
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(companies)
	}
}

// Find handler of a company
func (rh CompanyHTTPHandler) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		company, err := rh.Cs.Find(int64(id))
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(company)
	}
}

// Store handler of a company
func (rh CompanyHTTPHandler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var company model.Company
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &company)
		company.User = *middleware.GetUser(r.Context())
		company, err := rh.Cs.Store(company)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(company)
	}
}

// Update handler of a company
func (rh CompanyHTTPHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var company model.Company
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &company)
		company.ID = int64(id)
		company, err = rh.Cs.Update(company)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(company)
	}
}

// Destroy handler of a company
func (rh CompanyHTTPHandler) Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		company, err := rh.Cs.Destroy(int64(id))
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(company)
	}
}
