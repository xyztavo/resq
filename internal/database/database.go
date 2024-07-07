package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/xyztavo/resq/configs"
)

var db *sql.DB

// starts once the package loads
func init() {
	var err error
	// creates a db connection
	db, err = sql.Open("postgres", configs.GetDbConnectionString())
	if err != nil {
		log.Fatal(err)
	}
}

// reuses that db connection
func GetDb() *sql.DB {
	return db
}

// migrate db using the initalized db connection
func Migrate() error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(40) PRIMARY KEY,
	name VARCHAR(40) NOT NULL,
	role VARCHAR(40),
	email VARCHAR(40) UNIQUE, 
	password VARCHAR(200) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS companies (
	id VARCHAR(40) PRIMARY KEY,
	name VARCHAR(40) NOT NULL,
	description VARCHAR(40) NOT NULL, 
	rating DOUBLE PRECISION,
	creator_id VARCHAR(40) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS ngos (
	id VARCHAR(40) PRIMARY KEY,
	name VARCHAR(40) NOT NULL,
	description VARCHAR(40) NOT NULL, 
	rating DOUBLE PRECISION,
	creator_id VARCHAR(40) NOT NULL
	);
	`)
	if err != nil {
		return err
	}
	return nil
}
