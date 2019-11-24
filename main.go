package main

import (
	"database/sql"
	"net/http"
	"os"
	"log"

	"github.com/andrewesteves/tapagguapi/handler"
	"github.com/andrewesteves/tapagguapi/middleware"
	"github.com/andrewesteves/tapagguapi/repository"
	"github.com/andrewesteves/tapagguapi/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	mux := mux.NewRouter().StrictSlash(true)
	receiptRepository := repository.NewReceiptPostgresRepository(db)
	receiptService := service.NewReceiptService(receiptRepository)
	handler.NewReceiptHttpHandler(mux, receiptService)

	itemRepository := repository.NewItemPostgresRepository(db)
	itemService := service.ItemContractService(itemRepository)
	handler.NewItemHttpHandler(mux, itemService)

	mux.Use(middleware.CorsMiddleware{}.Enable)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.ListenAndServe(":" + port, mux)
}
