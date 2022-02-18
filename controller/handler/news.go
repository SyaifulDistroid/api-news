package handler

import (
	"api-news/model"
	Utils "api-news/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) NewsInsert(c *gin.Context) {
	var data model.News
	c.BindJSON(&data)
	News, err := h.usecase.InsertNews(&data)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Insert News", fmt.Sprintf("%v", err))
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Insert News", News)
}

func (h *Handler) NewsUpdate(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Error(err)
		return
	}

	var data model.News
	c.BindJSON(&data)
	News, err := h.usecase.UpdateNews(id, &data)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Update News", fmt.Sprintf("%v", err))
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Update News", News)
}

func (h *Handler) NewsDelete(c *gin.Context) {
	strID := c.Param("id")

	id, errID := strconv.Atoi(strID)
	if errID != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Error(errID)
		return
	}

	err := h.usecase.DeleteNews(id)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Delete News", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Delete News", nil)
}

func (h *Handler) NewsGetAll(c *gin.Context) {
	News, err := h.usecase.GetAllNews()

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show News or Data Not Found", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show News", News)
}

func (h *Handler) NewsGetByID(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Error(err)
		return
	}

	News, err := h.usecase.GetNewsByID(id)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show News or Data Not Found", fmt.Sprintf("%v", err))
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show News", News)
}

func (h *Handler) NewsGetByTopics(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Error(err)
		return
	}

	News, err := h.usecase.GetNewsByTopics(id)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show News or Data Not Found", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show News", News)
}

func (h *Handler) NewsGetByStatus(c *gin.Context) {
	status := c.Param("status")

	News, err := h.usecase.GetNewsByStatus(status)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show News or Data Not Found", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show News", News)

}
