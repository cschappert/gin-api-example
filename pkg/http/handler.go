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
