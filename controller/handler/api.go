package handler

import (
	"api-news/controller"
	"api-news/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase controller.UsecaseInterface
}

func CreateHandler(r *gin.Engine, usecase controller.UsecaseInterface) {
	Handler := &Handler{usecase}

	r.Use(middleware.RequestLoggerActivity())

	r.POST("/topics", Handler.TopicsInsert)
	r.PUT("/topics/:id", Handler.TopicsUpdate)
	r.DELETE("/topics/:id", Handler.TopicsDelete)
	r.GET("/topics", Handler.TopicsGetAll)
	r.GET("/topics/:id", Handler.TopicsGetByID)
	r.GET("/topic/:topic", Handler.TopicsGetByTopic)

	r.POST("/tags", Handler.TagsInsert)
	r.PUT("/tags/:id", Handler.TagsUpdate)
	r.DELETE("/tags/:id", Handler.TagsDelete)
	r.GET("/tags", Handler.TagsGetAll)
	r.GET("/tags/:id", Handler.TagsGetByID)
	r.GET("/tag/:tag", Handler.TagsGetByTag)

	r.POST("/news", Handler.NewsInsert)
	r.PUT("/news/:id", Handler.NewsUpdate)
	r.DELETE("/news/:id", Handler.NewsDelete)
	r.GET("/news", Handler.NewsGetAll)
	r.GET("/news/:id", Handler.NewsGetByID)
	r.GET("/news/topics/:id", Handler.NewsGetByTopics)
	r.GET("/news/status/:status", Handler.NewsGetByStatus)
}
