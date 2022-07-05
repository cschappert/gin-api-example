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

package http

import (
	"net/http"
	"strconv"

	api "github.com/cschappert/gin-api-example/pkg"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	AccountService api.AccountService
}

func NewHandler(g *gin.Engine, as api.AccountService) {
	h := &Handler{as}

	v1 := g.Group("/api/v1")

	accounts := v1.Group("/accounts")
	{
		accounts.GET("", h.ListAccounts)
		accounts.GET("/:id", h.GetAccount)
		accounts.PUT("", h.CreateAccount)
		accounts.DELETE("/:id", h.DeleteAccount)
	}

	/* TODO
	teams := v1.Group("/teams")
	{
		teams.GET("", h.ListTeams)
		teams.GET("/:id", h.GetTeam)
		teams.PUT("", h.CreateTeam)
		teams.DELETE("", h.DeleteTeam)
	}
	*/
}

func (h *Handler) ListAccounts(c *gin.Context) {
	res, err := h.AccountService.ListAccounts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := h.AccountService.GetAccount(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var account api.Account
	err := c.ShouldBindJSON(&account)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.AccountService.CreateAccount(&account)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func (h *Handler) DeleteAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.AccountService.DeleteAccount(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

/* TODO
func (h *Handler) ListTeams(c *gin.Context) {
	c.String(200, "Success")
}

func (h *Handler) GetTeam(c *gin.Context) {
	res := api.Team{
		Name:       "Bob",
		OwnerEmail: "bob@example.com",
	}
	c.JSON(200, res)
}

func (h *Handler) CreateTeam(c *gin.Context) {
	c.String(204, "Success")
}

func (h *Handler) DeleteTeam(c *gin.Context) {
	c.JSON(204, "Success")
}
*/
