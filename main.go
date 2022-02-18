package main

import (
	apiHandler "api-news/controller/handler"
	apiRepo "api-news/controller/repo"
	apiUsecase "api-news/controller/usecase"
	"api-news/model"
	"database/sql"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error(err)
	}
}

func setupLogOutput() {
	f, _ := os.Create("gin.Log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")
	port := viper.GetString("portlocal.port")
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("sslmode", "disable")
	connStr := fmt.Sprintf("%s?%s", connection, val.Encode())

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer db.Close()

	dbConn, err := gorm.Open("postgres", connStr)
	if err != nil {
		logrus.Error(err)
		return
	}

	err = dbConn.DB().Ping()
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println("Success Connect DB")

	defer func() {
		err = dbConn.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	dbConn.Debug().AutoMigrate(
		&model.Topics{},
		&model.News{},
		&model.Tags{},
	)

	setupLogOutput()
	router := gin.New()

	NewsRepo := apiRepo.CreateNewsRepoMysql(dbConn)
	TagsRepo := apiRepo.CreateTagsRepoMysql(dbConn)
	TopicsRepo := apiRepo.CreateTopicsRepoMysql(dbConn)

	Usecase := apiUsecase.CreateUsecase(NewsRepo, TopicsRepo, TagsRepo)
	apiHandler.CreateHandler(router, Usecase)

	fmt.Println("Starting Web Server at port : " + port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
