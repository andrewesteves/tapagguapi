package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/andrewesteves/tapagguapi/middleware"
	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/service"
	"github.com/andrewesteves/tapagguapi/transformer"
	"github.com/gorilla/mux"
)

// CategoryHTTPHandler struct
type CategoryHTTPHandler struct {
	Cs service.CategoryContractService
}

// NewCategoryHTTPHandler new category handler
func NewCategoryHTTPHandler(mux *mux.Router, categoryService service.CategoryContractService) {
	handler := &CategoryHTTPHandler{
		Cs: categoryService,
	}
	mux.HandleFunc("/categories/{id}", handler.Find()).Methods("GET")
	mux.HandleFunc("/categories/{id}", handler.Update()).Methods("PUT")
	mux.HandleFunc("/categories/{id}", handler.Destroy()).Methods("DELETE")
	mux.HandleFunc("/categories", handler.Store()).Methods("POST")
	mux.HandleFunc("/categories", handler.All()).Methods("GET")
}

// All handler of a categories
func (rh CategoryHTTPHandler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := rh.Cs.All(*middleware.GetUser(r.Context()))
		if err != nil {
			panic(err.Error())
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(categories)
	}
}

// Find handler of a category
func (rh CategoryHTTPHandler) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		category, err := rh.Cs.Find(int64(id))
		if err != nil {
			panic(err.Error())
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformer.CategoryTransformer{}.TransformOne(category))
	}
}

// Store handler of a category
func (rh CategoryHTTPHandler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category model.Category
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &category)
		category.User = *middleware.GetUser(r.Context())
		category, err := rh.Cs.Store(category)
		if err != nil {
			panic(err.Error())
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformer.CategoryTransformer{}.TransformOne(category))
	}
}

// Update handler of a category
func (rh CategoryHTTPHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category model.Category
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &category)
		category.ID = int64(id)
		category, err = rh.Cs.Update(category)
		if err != nil {
			panic(err.Error())
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformer.CategoryTransformer{}.TransformOne(category))
	}
}

// Destroy handler of a category
func (rh CategoryHTTPHandler) Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		category, err := rh.Cs.Destroy(int64(id))
		if err != nil {
			panic(err.Error())
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformer.CategoryTransformer{}.TransformOne(category))
	}
}
