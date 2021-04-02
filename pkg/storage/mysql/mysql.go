// Copyright 2021 Chris Schappert
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
