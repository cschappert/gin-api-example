package main

import (
	"github.com/cschappert/gin-api-example/pkg/http"
	"github.com/cschappert/gin-api-example/pkg/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	db := mysql.Open()
	defer db.Close()

	as := &mysql.AccountService{DB: db}
	r := gin.Default()
	http.NewHandler(r, as)

	r.Run()
}
