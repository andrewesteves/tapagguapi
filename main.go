package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/andrewesteves/tapagguapi/handler"
	"github.com/andrewesteves/tapagguapi/middleware"
	"github.com/andrewesteves/tapagguapi/repository"
	"github.com/andrewesteves/tapagguapi/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "4321"
	dbname   = "tapaggu"
	sslmode  = "disable"
)

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	file, err := ioutil.ReadFile("db.sql")
	if err != nil {
		panic(err.Error())
	}
	stmt := strings.Split(string(file), ";")
	for _, s := range stmt {
		_, err := db.Exec(s)
		if err != nil {
			panic(err.Error())
		}
	}

	mux := mux.NewRouter().StrictSlash(true)
	receiptRepository := repository.NewReceiptPostgresRepository(db)
	receiptService := service.NewReceiptService(receiptRepository)
	handler.NewReceiptHTTPHandler(mux, receiptService)

	itemRepository := repository.NewItemPostgresRepository(db)
	itemService := service.NewItemService(itemRepository)
	handler.NewItemHttpHandler(mux, itemService)

	userRepository := repository.NewUserPostgresRepository(db)
	userService := service.NewUserService(userRepository)
	handler.NewUserHttpHandler(mux, userService)

	auth := middleware.AuthMiddleware{Conn: db}
	mux.Use(middleware.CorsMiddleware{}.Enable)
	mux.Use(auth.Enable)
	http.ListenAndServe(":3000", mux)
}
