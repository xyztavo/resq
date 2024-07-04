package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func GetPort() string {
	return os.Getenv("PORT")
}

func GetDbConnectionString() string {
	var (
		DatabaseUser     = os.Getenv("DB_USER")
		DatabasePassword = os.Getenv("DB_PASSWORD")
		DatabaseName     = os.Getenv("DB_NAME")
	)
	return fmt.Sprintf("postgres://%v:%v@localhost:5432/%v?sslmode=disable", DatabaseUser, DatabasePassword, DatabaseName)
}
