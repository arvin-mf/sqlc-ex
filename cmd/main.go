package main

import (
	"log"
	"sj/config"
	"sj/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.SetupDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()

	handler.StartEngine(r, db)

	r.Run(":8000")
}
