package service

import (
	"github.com/edujudici/golang-crud-swagger/data/request"
	"github.com/edujudici/golang-crud-swagger/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
