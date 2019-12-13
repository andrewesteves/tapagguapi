package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/service"
	"github.com/gorilla/mux"
)

// UserHTTPHandler struct
type UserHTTPHandler struct {
	Us service.UserContractService
}

// NewUserHTTPHandler new user handler
func NewUserHTTPHandler(mux *mux.Router, userService service.UserContractService) {
	handler := &UserHTTPHandler{
		Us: userService,
	}
	mux.HandleFunc("/users/confirmation", handler.Confirmation()).Methods("GET")
	mux.HandleFunc("/users/recover", handler.Recover()).Methods("POST")
	mux.HandleFunc("/users/login", handler.Login()).Methods("POST")
	mux.HandleFunc("/users/logout", handler.Logout()).Methods("POST")
	mux.HandleFunc("/users/{id}", handler.Find()).Methods("GET")
	mux.HandleFunc("/users/{id}", handler.Update()).Methods("PUT")
	mux.HandleFunc("/users/{id}", handler.Destroy()).Methods("DELETE")
	mux.HandleFunc("/users", handler.Store()).Methods("POST")
	mux.HandleFunc("/users", handler.All()).Methods("GET")
}

// All handler of users
func (uh UserHTTPHandler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := uh.Us.All()
		if err != nil {
			panic(err.Error())
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}

// Find handler of user
func (uh UserHTTPHandler) Find() http.HandlerFunc {
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

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

// Store handler of user
func (uh UserHTTPHandler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)
		user, err := uh.Us.Store(user)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
			})
			return
		}
		err = service.Mailer{}.Send([]string{user.Email}, "welcome", []string{})
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
			})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

// Update handler of user
func (uh UserHTTPHandler) Update() http.HandlerFunc {
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

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

// Destroy handler of user
func (uh UserHTTPHandler) Destroy() http.HandlerFunc {
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

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

// Login handler of user
func (uh UserHTTPHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)
		u, err := uh.Us.Login(user)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
			})
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(u)
		}
	}
}

// Logout handler of user
func (uh UserHTTPHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)

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

// Recover password
func (uh UserHTTPHandler) Recover() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)
		dUser, err := uh.Us.FindBy("email", user.Email)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
			})
			return
		}
		err = service.Mailer{}.Send([]string{user.Email}, "recover", []string{dUser.Email, dUser.Token})
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
			})
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// Confirmation email
func (uh UserHTTPHandler) Confirmation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.ParseFiles("template/confirmation.html")
		if err != nil {
			panic(err.Error())
		}
		w.Header().Set("Content-Type", "text/html")
		tpl.Execute(w, nil)
	}
}
