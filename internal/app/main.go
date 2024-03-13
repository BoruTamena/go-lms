package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/BoruTamena/habitant_hub/config"
	"github.com/BoruTamena/habitant_hub/internal/app/handler"
	"github.com/BoruTamena/habitant_hub/internal/app/repository"
	"github.com/BoruTamena/habitant_hub/internal/app/service"
	_ "github.com/lib/pq"
)

func main() {
	db_config := *config.GetDbConnection()

	DbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", db_config.User, db_config.Password, db_config.Host, db_config.Port, db_config.DbName)

	fmt.Printf("start con on port :%v", DbUrl)
	con, err := sql.Open("postgres", DbUrl)

	if err != nil {
		log.Fatal(err)
	}

	// initializing repository

	rep := repository.NewUserRepository(con)
	ser := service.NewUserService(rep)

	userhandler := handler.NewUserHandler(ser)

	// Defining route
	http.HandleFunc("/user", userhandler.CreateUser)

	// creating Http server
	log.Fatal(http.ListenAndServe(":8000", nil))

}
