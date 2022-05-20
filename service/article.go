package service

import (
	"MedicalCare/model"
	"MedicalCare/pkg/e"
)

type ArticleService struct {
	Title   string `json:"title" form:"title"`
	Desc    string `json:"desc" form:"desc"`
	Content string `json:"content" form:"content"`
}

// Create 新建文章
func (service *ArticleService) Create() model.Response {
	code := e.Success

	return model.Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: "Successful write the article",
	}
}
