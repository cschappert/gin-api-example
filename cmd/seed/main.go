// Copyright 2022 Chris Schappert
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

package main

import (
	"log"

	storage "github.com/cschappert/gin-api-example/pkg/storage/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:pw@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=UTC"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	db.Create(&storage.Account{
		Name:  "Bob",
		Email: "bob@example.com",
		Team: storage.Team{
			Name: "team1",
		},
	})

	db.Create(&storage.Account{
		Name:  "Alice",
		Email: "alice@example.com",
		Team: storage.Team{
			Name: "team2",
		},
	})

}
