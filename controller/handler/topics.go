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

func (h *Handler) TopicsInsert(c *gin.Context) {
	var data model.Topics
	c.BindJSON(&data)
	Topics, err := h.usecase.InsertTopics(&data)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Insert Topics", fmt.Sprintf("%v", err))
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Insert Topics", Topics)
}

func (h *Handler) TopicsUpdate(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", fmt.Sprintf("%v", err))
		logrus.Error(err)
		return
	}

	var data model.Topics
	c.BindJSON(&data)
	Topics, err := h.usecase.UpdateTopics(id, &data)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Update Topics", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Update Topics", Topics)
}

func (h *Handler) TopicsDelete(c *gin.Context) {
	strID := c.Param("id")

	id, errID := strconv.Atoi(strID)
	if errID != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Error(errID)
		return
	}

	err := h.usecase.DeleteTopics(id)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Delete Topics", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Delete Topics", nil)
}

func (h *Handler) TopicsGetAll(c *gin.Context) {
	Topics, err := h.usecase.GetAllTopics()

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show Topics or Data Not Found", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show Topics", Topics)
}

func (h *Handler) TopicsGetByID(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Error(err)
		return
	}

	Topics, err := h.usecase.GetTopicsByID(id)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show Topics or Data Not Found", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show Topics", Topics)
}

func (h *Handler) TopicsGetByTopic(c *gin.Context) {
	topic := c.Param("topic")

	Topics, err := h.usecase.GetTopicsByTopic(topic)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show Topics or Data Not Found", nil)
		logrus.Error(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show Topics", Topics)
}
