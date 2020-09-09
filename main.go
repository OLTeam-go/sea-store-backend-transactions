package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	bankRepo "github.com/OLTeam-go/sea-store-backend-transactions/bank/repository/postgresql"
	bankUsecase "github.com/OLTeam-go/sea-store-backend-transactions/bank/usecase"
	cartRepo "github.com/OLTeam-go/sea-store-backend-transactions/cart/repository/postgresql"
	cartUsecase "github.com/OLTeam-go/sea-store-backend-transactions/cart/usecase"
	cartItemRepo "github.com/OLTeam-go/sea-store-backend-transactions/cart_item/repository/postgresql"
	cartItemUsecase "github.com/OLTeam-go/sea-store-backend-transactions/cart_item/usecase"
	database "github.com/OLTeam-go/sea-store-backend-transactions/db"
	dTransactions "github.com/OLTeam-go/sea-store-backend-transactions/delivery/http"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Service Transaction API
// @version 1.0
// @description Api Documentation for Service Transaction

// @contact.name OLTeamgo API Support
// @contact.email yoganandamahaputra@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host sea-store-backend-transaction.herokuapp.com
// @BasePath /api

func migrations(url string) {
	fmt.Println("starting migrations")
	m, err := migrate.New(
		"file://db/migrations",
		url)
	if err != nil {
		log.Println(err.Error())
	}
	if err := m.Up(); err != nil {
		log.Println(err.Error())
	}
	fmt.Println("migrations done")
}

func main() {
	_ = godotenv.Load()

	dbURL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		panic("DATABASE_URL did not exists")
	}
	migrations(dbURL)

	db, err := database.GetInstance()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	pagesize, err := strconv.Atoi(os.Getenv("PAGESIZE"))
	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		log.Println(err.Error())
	}
	port := os.Getenv("PORT")

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	tc := time.Duration(timeout) * time.Second

	bRepo := bankRepo.New(db)
	bUsecase := bankUsecase.New(bRepo, tc)

	cRepo := cartRepo.New(db, pagesize)
	cUsecase := cartUsecase.New(cRepo, tc)

	cItRepo := cartItemRepo.New(db, pagesize)
	cItUsecase := cartItemUsecase.New(cItRepo, cRepo, tc)

	dTransactions.New(e, bUsecase, cUsecase, cItUsecase)

	log.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
