package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type User struct{
	Email string `json:"email"`
	Password int `json:"password"`
}

type Config struct{
	DBProperties struct{
		Username      string `json:"user"`
		Password      string `json:"password"`
		Port          string `json:"port"`
		Database_name string `json:"database_name"`
		Address       string `json:"address"`
	} `yaml:"database"`
}
func GetDatabase() (*sql.DB, error) {

	db, err := sql.Open("mysql", root : Hanuman1998@tcp("localhost:3306")/mydb
	if err != nil {
		log.Panic(err.Error())
	}
	log.Println( "DB Connection Successful")
	return db, nil
}
