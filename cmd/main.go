package main

import (
	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/server"
	"fmt"
	"log"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	// Init echo and .env
	runEcho := echo.New()
	port := config.GetAppPort()

	// Setup Middleware
	runEcho.Use(middleware.RequestLogger())
	runEcho.Use(middleware.Recover())

	// Open Connection to Database
	db, err := database.ConnectPostgres()
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		if errClose := sqlDB.Close(); errClose != nil {
			fmt.Println(errClose)
		}
	}()

	// Registering Route
	server.Register(runEcho, db)
	if err := runEcho.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
