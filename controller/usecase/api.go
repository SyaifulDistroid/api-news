package usecase

import "api-news/controller"

type UsecaseImpl struct {
	NewsDB   controller.News
	TopicsDB controller.Topics
	TagsDB   controller.Tags
}

func CreateUsecase(NewsRepo controller.News, TopicsRepo controller.Topics, TagsRepo controller.Tags) controller.UsecaseInterface {
	return &UsecaseImpl{NewsDB: NewsRepo, TopicsDB: TopicsRepo, TagsDB: TagsRepo}
}
