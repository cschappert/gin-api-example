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
	"github.com/cschappert/gin-api-example/pkg/http"
	"github.com/cschappert/gin-api-example/pkg/storage/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	db := mysql.Open()

	// Create a mysql implementation of the AccountService (a mock could be used here instead if unit testing)
	as := &mysql.AccountService{DB: db}
	r := gin.Default()

	// Pass the mysql implementation of the AccountService to the handler
	http.NewHandler(r, as)

	_ = r.Run()
}
