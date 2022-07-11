package main

import (
	"fmt"
	"golang-sample-injection/config"
	"golang-sample-injection/model"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.NewConfig()
	db := cfg.DbConn()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	routerEngine := gin.Default()
	routerGroup := routerEngine.Group("/api")
	routerGroup.POST("/auth/login", func(ctx *gin.Context) {
		var login model.Login
		if err := ctx.ShouldBindJSON(&login); err != nil {
			ctx.JSONP(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		var userCred model.UserCredential
		//sql := fmt.Sprintf("SELECT * FROM user_credential WHERE user_name='%s' and user_password='%s'",
		//	login.User, login.Password,
		//)
		// agar tidak terkena sql injection menggunakan query param seperti $1 atau ?
		// namanya sanitasi
		sql := "SELECT * FROM user_credential WHERE user_name=$1 and user_password=$2 and is_blocked='f'"
		log.Println("sql", sql)

		err := db.Get(&userCred, sql, login.User, login.Password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"token":   "123",
			"message": userCred,
			//"message": "PONG",
		})
	})
	var apiPort = config.ApiConfig{}
	apiPort.ApiHost = os.Getenv("API_HOST")
	apiPort.ApiPort = os.Getenv("API_PORT")
	listenAddress := fmt.Sprintf("%s:%s", apiPort.ApiHost, apiPort.ApiPort)
	//listenAddress := fmt.Sprintf("%s:%s", "localhost", "8888")

	err := routerEngine.Run(listenAddress)
	if err != nil {
		panic(err)
	}
}

/*
set DB_DRIVER=postgres
set DB_HOST=localhost
set DB_PORT=5432
set DB_USER=postgres
set DB_PASSWORD=87654321
set DB_NAME=db_credential
set API_HOST=localhost
set API_PORT=8888

postman --> localhost:8888/api/auth/login
{
    "userName": "alien' or 'what' = 'what",
    "userPassword":"alien' or 'what' = 'what"
}
*/
