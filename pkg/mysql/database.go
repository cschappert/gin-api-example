package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var _db *sql.DB

func Open() *sql.DB {
	if _db == nil {
		host := "127.0.0.1"
		port := "13306"
		user := "user"
		password := "password"
		database := "database"
		maxLifeTime := 60

		url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4", user, password, host, port, database)

		db, err := sql.Open("mysql", url)
		if err != nil {
			log.Fatal("DB ERROR: ", err)
		}

		db.SetMaxIdleConns(100)
		db.SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Second)

		err = db.Ping()
		if err != nil {
			log.Fatal("DB ERROR: ", err)
		}

		_db = db
	}

	return _db
}
