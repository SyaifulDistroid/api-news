package usecase

import (
	"api-news/model"
	"errors"
)

func (u *UsecaseImpl) InsertTags(req *model.Tags) (*model.Tags, error) {
	_, err := u.GetTagsByTag(req.Tag)
	if err == nil {
		return nil, errors.New("Tag Already")
	}
	return u.TagsDB.InsertTags(req)
}

func (u *UsecaseImpl) UpdateTags(id int, req *model.Tags) (*model.Tags, error) {
	_, err := u.GetTagsByTag(req.Tag)
	if err == nil {
		return nil, errors.New("Tag Already")
	}
	return u.TagsDB.UpdateTags(id, req)
}

func (u *UsecaseImpl) DeleteTags(id int) error {
	return u.TagsDB.DeleteTags(id)
}

func (u *UsecaseImpl) GetAllTags() ([]*model.Tags, error) {
	return u.TagsDB.GetAllTags()
}

func (u *UsecaseImpl) GetTagsByID(id int) (*model.Tags, error) {
	return u.TagsDB.GetTagsByID(id)
}

func (u *UsecaseImpl) GetTagsByTag(tag string) (*model.Tags, error) {
	return u.TagsDB.GetTagsByTag(tag)
}
