package main

import (
	"database/sql"
	"net/http"

	"github.com/andrewesteves/tapagguapi/handler"
	"github.com/andrewesteves/tapagguapi/repository"
	"github.com/andrewesteves/tapagguapi/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=root password=4321 dbname=tapaggu sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	mux := mux.NewRouter().StrictSlash(true)
	receiptRepository := repository.NewReceiptPostgresRepository(db)
	receiptService := service.NewReceiptService(receiptRepository)
	handler.NewReceiptHttpHandler(mux, receiptService)

	http.ListenAndServe(":3000", mux)
}
