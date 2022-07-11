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
	api "github.com/cschappert/gin-api-example/pkg"
	"github.com/cschappert/gin-api-example/pkg/http"
	storage "github.com/cschappert/gin-api-example/pkg/storage/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	// get the DB
	db := storage.Open()

	// create a new storage.AccountRepository, passing in the DB
	ar := storage.NewAccountRepository(db)

	// Note: api.NewAccountService takes an api.AccountRepository, which is an interface.
	// When passing an implementation (in this case, storage.AccountRepository) we have to pass a *pointer*
	// to the implementation.
	as := api.NewAccountService(&ar)
	r := gin.Default()

	// Pass the mysql implementation of the AccountService to the handler
	http.NewHandler(r, as)

	_ = r.Run()
}
