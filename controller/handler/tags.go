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

func (h *Handler) TagsInsert(c *gin.Context) {
	var data model.Tags
	c.BindJSON(&data)
	Tags, err := h.usecase.InsertTags(&data)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Insert Tags", fmt.Sprintf("%v", err))
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Insert Tags", Tags)
}

func (h *Handler) TagsUpdate(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", fmt.Sprintf("%v", err))
		logrus.Error(err)
		return
	}

	var data model.Tags
	c.BindJSON(&data)
	Tags, err := h.usecase.UpdateTags(id, &data)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Update Tags", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Update Tags", Tags)
}

func (h *Handler) TagsDelete(c *gin.Context) {
	strID := c.Param("id")

	id, errID := strconv.Atoi(strID)
	if errID != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Error(errID)
		return
	}

	err := h.usecase.DeleteTags(id)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Delete Tags", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Delete Tags", nil)
}

func (h *Handler) TagsGetAll(c *gin.Context) {
	Tags, err := h.usecase.GetAllTags()

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show Tags or Data Not Found", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show Tags", Tags)
}

func (h *Handler) TagsGetByID(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Error(err)
		return
	}

	Tags, err := h.usecase.GetTagsByID(id)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show Tags or Data Not Found", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show Tags", Tags)
}

func (h *Handler) TagsGetByTag(c *gin.Context) {
	Tag := c.Param("tag")

	Tags, err := h.usecase.GetTagsByTag(Tag)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show Tags or Data Not Found", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show Tags", Tags)
}
