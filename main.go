package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pradiptadh/majoo/pkg/config"
	"github.com/pradiptadh/majoo/pkg/db"
	"github.com/pradiptadh/majoo/pkg/routes"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// set time zone
	os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	dbUrl := config.LoadDBConfig()
	h := db.Init(dbUrl)
	r := routes.SetupRouter(h)
	r.Run()
}
