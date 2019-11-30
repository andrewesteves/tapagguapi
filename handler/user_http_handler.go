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

type UserHttpHandler struct {
	Us service.UserContractService
}

func NewUserHttpHandler(mux *mux.Router, userService service.UserContractService) {
	handler := &UserHttpHandler{
		Us: userService,
	}
	mux.HandleFunc("/users/login", handler.Login()).Methods("POST")
	mux.HandleFunc("/users/logout", handler.Logout()).Methods("POST")
	mux.HandleFunc("/users/{id}", handler.Find()).Methods("GET")
	mux.HandleFunc("/users/{id}", handler.Update()).Methods("PUT")
	mux.HandleFunc("/users/{id}", handler.Destroy()).Methods("DELETE")
	mux.HandleFunc("/users", handler.Store()).Methods("POST")
	mux.HandleFunc("/users", handler.All()).Methods("GET")
}

func (uh UserHttpHandler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := uh.Us.All()
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}

func (uh UserHttpHandler) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		user, err := uh.Us.Find(int64(id))
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

func (uh UserHttpHandler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)
		user, err := uh.Us.Store(user)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		user.Password = ""
		json.NewEncoder(w).Encode(user)
	}
}

func (uh UserHttpHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)
		user.ID = int64(id)
		user, err = uh.Us.Update(user)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

func (uh UserHttpHandler) Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err.Error())
		}
		user, err := uh.Us.Destroy(int64(id))
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

func (uh UserHttpHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)
		w.Header().Add("Content-Type", "application/json")
		u, err := uh.Us.Login(user)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(nil)
		} else {
			u.Password = ""
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(u)
		}
	}
}

func (uh UserHttpHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)
		w.Header().Add("Content-Type", "application/json")
		_, err := uh.Us.Logout(user)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(nil)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(model.User{})
		}
	}
}