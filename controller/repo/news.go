package repo

import (
	"api-news/controller"
	"api-news/model"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type news struct {
	Conn *gorm.DB
}

func CreateNewsRepoMysql(DB *gorm.DB) controller.News {
	return &news{DB}
}

func (n *news) InsertNews(req *model.News) (*model.News, error) {
	if err := n.Conn.Table("news_tb").Save(&req).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Insert News is Failed")
	}
	return req, nil
}

func (n *news) UpdateNews(id int, req *model.News) (*model.News, error) {
	news := new(model.News)

	if err := n.Conn.Table("news_tb").Where("id=?", id).First(&news).Update(&req).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Update News is Failed")
	}
	return news, nil
}

func (n *news) DeleteNews(id int) error {
	if err := n.Conn.Table("news_tb").Where("id = ?", id).Delete(&model.News{}).Error; err != nil {
		logrus.Error(err)
		return errors.New("Delete News Failed")
	}
	return nil
}

func (n *news) GetAllNews() ([]*model.News, error) {
	newsList := make([]*model.News, 0)

	if err := n.Conn.Table("news_tb").Find(&newsList).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Get All News Failed")
	}

	return newsList, nil
}

func (n *news) GetNewsByID(id int) (*model.News, error) {
	news := new(model.News)

	if err := n.Conn.Table("news_tb").Where("id = ?", id).First(&news).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Get News By ID Failed")
	}
	return news, nil
}

func (n *news) GetNewsByTopics(id int) ([]*model.News, error) {
	newsList := make([]*model.News, 0)

	if err := n.Conn.Table("news_tb").Where("topics_id = ?", id).Find(&newsList).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Get News By ID Failed")
	}
	return newsList, nil
}

func (n *news) GetNewsByStatus(status string) ([]*model.News, error) {
	newsList := make([]*model.News, 0)

	if err := n.Conn.Table("news_tb").Where("status = ?", status).Find(&newsList).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Get News By ID Failed")
	}
	return newsList, nil
}
