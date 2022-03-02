package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

var DB *sql.DB

func InitDB() {
	viper.SetConfigName(".env")
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	user := viper.GetString("DATABASE_USER")
	password := viper.GetString("DATABASE_PASSWORD")
	dbname := viper.GetString("DATABASE_NAME")
	port := viper.GetString("DATABASE_PORT")
	host := viper.GetString("DATABASE_HOST")
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	//fmt.Println(connStr)
	var err error
	DB, err = sql.Open("postgres",
		connStr)
	if err != nil {
		log.Fatal(err)
	}
	err1 := DB.Ping()
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("Connecting to db...")
	//defer DB.Close()
}
