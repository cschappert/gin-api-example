package main

import (
	"log"

	"github.com/cschappert/gin-api-example/pkg/storage/mysql"
	mysqlgorm "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:pw@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=UTC"
	db, err := gorm.Open(mysqlgorm.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	db.Create(&mysql.Account{
		Name:  "Bob",
		Email: "bob@example.com",
		Team: mysql.Team{
			Name: "team",
		},
	})

}
