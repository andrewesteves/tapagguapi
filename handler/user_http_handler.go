package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/andrewesteves/tapagguapi/config"
	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/service"
	"github.com/andrewesteves/tapagguapi/transformer"
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
	mux.HandleFunc("/users/resend", handler.Resend()).Methods("POST")
	mux.HandleFunc("/users/reset_confirmation", handler.ResetConfirmation()).Methods("GET")
	mux.HandleFunc("/users/reset", handler.Reset()).Methods("POST")
	mux.HandleFunc("/users/new_password", handler.NewPassword()).Methods("GET")
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
		err = service.Mailer{}.Send([]string{user.Email}, "welcome", []string{user.Email, user.Remember})
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
			})
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": config.LangConfig{}.I18n()["welcome"],
		})
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
		user, err = uh.Us.Update(user, user.RenewPassword != "")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformer.UserTransformer{}.TransformOne(user))
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
		json.NewEncoder(w).Encode(transformer.UserTransformer{}.TransformOne(user))
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
			return
		}
		if u.Active == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": config.LangConfig{}.I18n()["user_inactive"],
			})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformer.UserTransformer{}.TransformOne(u))
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
		err = service.Mailer{}.Send([]string{user.Email}, "recover", []string{dUser.Email, dUser.Remember})
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
		if r.URL.Query().Get("email") != "" && r.URL.Query().Get("token") != "" {
			user, err := uh.Us.FindByArgs(map[string]interface{}{
				"email":    r.URL.Query().Get("email"),
				"remember": r.URL.Query().Get("token"),
			})
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{
					"message": err.Error(),
				})
				return
			}
			user.Active = 1
			_, err = uh.Us.Update(user, false)
			var tpl *template.Template
			tpl, err = template.ParseFiles("template/layout.html", "template/confirmation.html")
			if err != nil {
				panic(err.Error())
			}
			w.Header().Set("Content-Type", "text/html")
			tpl.ExecuteTemplate(w, "layout", nil)
		}
	}
}

// NewPassword email
func (uh UserHTTPHandler) NewPassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("email") != "" && r.URL.Query().Get("token") != "" {
			var tpl *template.Template
			tpl, err := template.ParseFiles("template/layout.html", "template/new_password.html")
			if err != nil {
				panic(err.Error())
			}
			w.Header().Set("Content-Type", "text/html")
			tpl.ExecuteTemplate(w, "layout", struct {
				Email  string
				Token  string
				Action string
			}{
				r.URL.Query().Get("email"),
				r.URL.Query().Get("token"),
				config.EnvConfig{}.App.URL + "/users/reset",
			})
		}
	}
}

// Reset password
func (uh UserHTTPHandler) Reset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("email") != "" && r.FormValue("token") != "" {
			user, err := uh.Us.FindByArgs(map[string]interface{}{
				"email":    r.FormValue("email"),
				"remember": r.FormValue("token"),
			})
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{
					"message": err.Error(),
				})
				return
			}
			user.Password = r.FormValue("password")
			user.Remember = r.FormValue("token")
			_, err = uh.Us.UpdateRecover(user)
			http.Redirect(w, r, config.EnvConfig{}.App.URL+"/users/reset_confirmation", http.StatusSeeOther)
		}
	}
}

// ResetConfirmation password
func (uh UserHTTPHandler) ResetConfirmation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tpl *template.Template
		tpl, err := template.ParseFiles("template/layout.html", "template/reset_confirmation.html")
		if err != nil {
			panic(err.Error())
		}
		w.Header().Set("Content-Type", "text/html")
		tpl.ExecuteTemplate(w, "layout", nil)
	}
}

// Resend welcome email
func (uh UserHTTPHandler) Resend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)
		user, err := uh.Us.FindBy("email", user.Email)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": config.LangConfig{}.I18n()["email_invalid"],
			})
			return
		}
		if user.Active == 0 {
			err = service.Mailer{}.Send([]string{user.Email}, "welcome", []string{user.Email, user.Remember})
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusServiceUnavailable)
				json.NewEncoder(w).Encode(map[string]string{
					"message": err.Error(),
				})
				return
			}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": config.LangConfig{}.I18n()["email_send"],
		})
	}
}
