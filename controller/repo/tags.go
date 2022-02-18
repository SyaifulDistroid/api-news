package repo

import (
	"api-news/controller"
	"api-news/model"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type tags struct {
	Conn *gorm.DB
}

func CreateTagsRepoMysql(DB *gorm.DB) controller.Tags {
	return &tags{DB}
}

func (t *tags) InsertTags(req *model.Tags) (*model.Tags, error) {
	if err := t.Conn.Table("tags_tb").Save(&req).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Insert Tags is Failed")
	}
	return req, nil
}

func (t *tags) UpdateTags(id int, req *model.Tags) (*model.Tags, error) {
	tags := new(model.Tags)

	if err := t.Conn.Table("tags_tb").Where("id=?", id).First(&tags).Update(&req).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Update Tags is Failed")
	}
	return tags, nil
}

func (t *tags) DeleteTags(id int) error {
	if err := t.Conn.Table("tags_tb").Where("id = ?", id).Delete(&model.Tags{}).Error; err != nil {
		logrus.Error(err)
		return errors.New("Delete Tags Failed")
	}
	return nil
}

func (t *tags) GetAllTags() ([]*model.Tags, error) {
	tagsList := make([]*model.Tags, 0)

	if err := t.Conn.Table("tags_tb").Find(&tagsList).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Get All Tags Failed")
	}

	return tagsList, nil
}

func (t *tags) GetTagsByID(id int) (*model.Tags, error) {
	tags := new(model.Tags)

	if err := t.Conn.Table("tags_tb").Where("id = ?", id).First(&tags).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Get Tags By ID Failed")
	}
	return tags, nil
}

func (t *tags) GetTagsByTag(tag string) (*model.Tags, error) {
	tags := new(model.Tags)

	if err := t.Conn.Table("tags_tb").Where("tag = ?", tag).First(&tags).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Get Tags By Tag Failed")
	}
	return tags, nil
}
