package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect(){
	var err error
	host:=os.Getenv("DB_HOST")
	port:=os.Getenv("DB_PORT")
	username:=os.Getenv("DB_USER")
	password:=os.Getenv("DB_PASSWORD")
	dbName:=os.Getenv("DB_NAME")
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}else{
		fmt.Println("Database connected successfully!")
	}
}