package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/xyztavo/resq/configs"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", configs.GetDbConnectionString())
	if err != nil {
		log.Fatal(err)
	}
}

func Migrate() error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS company_admin (
	id UUID PRIMARY KEY,
	name VARCHAR(40),
	email VARCHAR(40) UNIQUE, 
	password VARCHAR(200)
	);`)
	if err != nil {
		return err
	}
	return nil
}
