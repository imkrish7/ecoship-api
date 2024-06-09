package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/imkrish7/ecoship-api/application"
	"github.com/imkrish7/ecoship-api/config"
	"github.com/imkrish7/ecoship-api/db"
	"github.com/imkrish7/ecoship-api/env"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	var config config.Config

	config.Port = env.Get("PORT", "3000")
	config.Addr = env.Get("ADDRESS", "http://localhost:3000")
	config.Db.Database = env.Get("DB_DATABASE", "eco")
	config.Db.Host = env.Get("DB_HOST", "localhost")
	config.Db.Password = env.Get("DB_PASSWORD", "eco123")
	config.Db.Port = env.Get("DB_PORT", "5435")
	config.Db.User = env.Get("DB_USER", "ecouser")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Db.Host, config.Db.User, config.Db.Password, config.Db.Database, config.Db.Port)
	log.Print(dsn)
	dbinstance, err := db.New(dsn)

	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	} else {
		log.Println("Data base connected")
	}

	app := &application.Application{
		DB:     dbinstance,
		CONFIG: config,
	}

	// http.ListenAndServe(":3000", h2c.NewHandler(AppRoutes(app), &http2.Server{}))

	if err := http.ListenAndServe(":4567", h2c.NewHandler(AppRoutes(app), &http2.Server{})); err != nil {
		panic(err)
	} else {
		log.Println("Server started....")
	}

}
