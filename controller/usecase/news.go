package usecase

import (
	"api-news/model"
	"errors"
	"strings"
)

func (u *UsecaseImpl) InsertNews(req *model.News) (*model.News, error) {
	_, err := u.TopicsDB.GetTopicsByID(req.TopicsID)
	if err != nil {
		return nil, errors.New("Topic Not Found")
	}

	for i := 0; i < len(req.Tags); i++ {
		_, err := u.TagsDB.GetTagsByTag(req.Tags[i])
		if err != nil {
			return nil, errors.New("Tag Not Found")
		}
	}

	stats := strings.ToLower(req.Status)
	if stats != "DRAFT" && stats != "DELETED" && stats != "PUBLISH" {
		return nil, errors.New("Status Only 'DRAFT', 'DELETED' OR 'PUBLISH'")
	}
	req.Status = stats

	return u.NewsDB.InsertNews(req)
}

func (u *UsecaseImpl) UpdateNews(id int, req *model.News) (*model.News, error) {
	_, err := u.TopicsDB.GetTopicsByID(req.TopicsID)
	if err != nil {
		return nil, errors.New("Topic Not Found")
	}

	for i := 0; i < len(req.Tags); i++ {
		_, err := u.TagsDB.GetTagsByTag(req.Tags[i])
		if err != nil {
			return nil, errors.New("Tag Not Found")
		}
	}

	stats := strings.ToLower(req.Status)
	if stats != "DRAFT" && stats != "DELETED" && stats != "PUBLISH" {
		return nil, errors.New("Status Only 'DRAFT', 'DELETED' OR 'PUBLISH'")
	}
	req.Status = stats

	return u.NewsDB.UpdateNews(id, req)
}

func (u *UsecaseImpl) DeleteNews(id int) error {
	return u.NewsDB.DeleteNews(id)
}

func (u *UsecaseImpl) GetAllNews() ([]*model.News, error) {
	return u.NewsDB.GetAllNews()
}

func (u *UsecaseImpl) GetNewsByID(id int) (*model.News, error) {
	return u.NewsDB.GetNewsByID(id)
}

func (u *UsecaseImpl) GetNewsByTopics(id int) ([]*model.News, error) {
	return u.NewsDB.GetNewsByTopics(id)
}

func (u *UsecaseImpl) GetNewsByStatus(status string) ([]*model.News, error) {
	status = strings.ToLower(status)
	return u.NewsDB.GetNewsByStatus(status)
}
