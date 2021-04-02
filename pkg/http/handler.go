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

package http

import (
	api "github.com/cschappert/gin-api-example/pkg"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	AccountService api.AccountService
}

func (h *Handler) ListAccounts(c *gin.Context) {
	c.String(200, "Success")
}

func (h *Handler) GetAccount(c *gin.Context) {
	res := api.Account{
		Name: "Bob",
	}
	c.JSON(200, res)
}

func NewHandler(g *gin.Engine, as api.AccountService) {
	h := &Handler{as}

	v1 := g.Group("/api/v1")

	accounts := v1.Group("/accounts")
	{
		accounts.GET("", h.ListAccounts)
		accounts.GET("/:id", h.GetAccount)
		//accounts.PUT("", h.CreateAccount)
		//accounts.DELETE("", h.DeleteAccount)
	}
}
