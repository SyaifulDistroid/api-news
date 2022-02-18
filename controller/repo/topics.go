package repo

import (
	"api-news/controller"
	"api-news/model"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type topics struct {
	Conn *gorm.DB
}

func CreateTopicsRepoMysql(DB *gorm.DB) controller.Topics {
	return &topics{DB}
}

func (t *topics) InsertTopics(req *model.Topics) (*model.Topics, error) {
	if err := t.Conn.Table("topics_tb").Save(&req).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Insert Topics is Failed")
	}
	return req, nil
}

func (t *topics) UpdateTopics(id int, req *model.Topics) (*model.Topics, error) {
	topics := new(model.Topics)

	if err := t.Conn.Table("topics_tb").Where("id=?", id).First(&topics).Update(&req).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Update Topics is Failed")
	}
	return topics, nil
}

func (t *topics) DeleteTopics(id int) error {
	if err := t.Conn.Table("topics_tb").Where("id = ?", id).Delete(&model.Topics{}).Error; err != nil {
		logrus.Error(err)
		return errors.New("Delete Topics Failed")
	}
	return nil
}

func (t *topics) GetAllTopics() ([]*model.Topics, error) {
	topicsList := make([]*model.Topics, 0)

	if err := t.Conn.Table("topics_tb").Find(&topicsList).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Get All Topics Failed")
	}

	return topicsList, nil
}

func (t *topics) GetTopicsByID(id int) (*model.Topics, error) {
	topics := new(model.Topics)

	if err := t.Conn.Table("topics_tb").Where("id = ?", id).First(&topics).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Get Topics By ID Failed")
	}
	return topics, nil
}

func (t *topics) GetTopicsByTopic(topic string) (*model.Topics, error) {
	topics := new(model.Topics)

	if err := t.Conn.Table("topics_tb").Where("topics = ?", topic).First(&topics).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Get Topics By Topics Failed")
	}
	return topics, nil
}
