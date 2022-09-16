package config

import (
	"fmt"
	"os"
)

func LoadDBConfig() (dbURL string) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/majoo?parseTime=true&loc=Local", username, password, host, port)
	return dsn
}
