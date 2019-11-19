package main

import (
	"database/sql"
	"net/http"

	"github.com/andrewesteves/tapagguapi/handler"
	"github.com/andrewesteves/tapagguapi/repository"
	"github.com/andrewesteves/tapagguapi/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "root:4321@tcp(127.0.0.1:3306)/tapaggu?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	mux := mux.NewRouter().StrictSlash(true)
	receiptRepository := repository.NewReceiptMySQLRepository(db)
	receiptService := service.NewReceiptService(receiptRepository)
	handler.NewReceiptHttpHandler(mux, receiptService)

	http.ListenAndServe(":3000", mux)
}
