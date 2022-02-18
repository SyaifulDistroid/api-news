package controller

import "api-news/model"

type News interface {
	InsertNews(req *model.News) (*model.News, error)
	UpdateNews(id int, req *model.News) (*model.News, error)
	DeleteNews(id int) error
	GetAllNews() ([]*model.News, error)
	GetNewsByID(id int) (*model.News, error)
	GetNewsByTopics(id int) ([]*model.News, error)
	GetNewsByStatus(status string) ([]*model.News, error)
}

type Topics interface {
	InsertTopics(req *model.Topics) (*model.Topics, error)
	UpdateTopics(id int, req *model.Topics) (*model.Topics, error)
	DeleteTopics(id int) error
	GetAllTopics() ([]*model.Topics, error)
	GetTopicsByID(id int) (*model.Topics, error)
	GetTopicsByTopic(topic string) (*model.Topics, error)
}

type Tags interface {
	InsertTags(req *model.Tags) (*model.Tags, error)
	UpdateTags(id int, req *model.Tags) (*model.Tags, error)
	DeleteTags(id int) error
	GetAllTags() ([]*model.Tags, error)
	GetTagsByID(id int) (*model.Tags, error)
	GetTagsByTag(tag string) (*model.Tags, error)
}
