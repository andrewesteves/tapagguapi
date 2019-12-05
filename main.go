package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
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
	handler.NewItemHTTPHandler(mux, itemService)

	userRepository := repository.NewUserPostgresRepository(db)
	userService := service.NewUserService(userRepository)
	handler.NewUserHTTPHandler(mux, userService)

	categoryRepository := repository.NewCategoryPostgresRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	handler.NewCategoryHTTPHandler(mux, categoryService)

	companyRepository := repository.NewCompanyPostgresRepository(db)
	companyService := service.NewCompanyService(companyRepository)
	handler.NewCompanyHTTPHandler(mux, companyService)

	auth := middleware.AuthMiddleware{Conn: db}
	mux.Use(middleware.CorsMiddleware{}.Enable)
	mux.Use(auth.Enable)
	mux.Use(middleware.ContentMiddleware{}.Enable)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set")
	}
	log.Println(http.ListenAndServe(":"+port, mux))
}
