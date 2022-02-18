package usecase

import (
	"api-news/model"
	"errors"
)

func (u *UsecaseImpl) InsertTopics(req *model.Topics) (*model.Topics, error) {
	_, err := u.GetTopicsByTopic(req.Topics)
	if err == nil {
		return nil, errors.New("Topics Already")
	}
	return u.TopicsDB.InsertTopics(req)
}

func (u *UsecaseImpl) UpdateTopics(id int, req *model.Topics) (*model.Topics, error) {
	_, err := u.GetTopicsByTopic(req.Topics)
	if err == nil {
		return nil, errors.New("Topics Already")
	}
	return u.TopicsDB.UpdateTopics(id, req)
}

func (u *UsecaseImpl) DeleteTopics(id int) error {
	return u.TopicsDB.DeleteTopics(id)
}

func (u *UsecaseImpl) GetAllTopics() ([]*model.Topics, error) {
	return u.TopicsDB.GetAllTopics()
}

func (u *UsecaseImpl) GetTopicsByID(id int) (*model.Topics, error) {
	return u.TopicsDB.GetTopicsByID(id)
}

func (u *UsecaseImpl) GetTopicsByTopic(topic string) (*model.Topics, error) {
	return u.TopicsDB.GetTopicsByTopic(topic)
}
