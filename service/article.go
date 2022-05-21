package service

import (
	"MedicalCare/model"
	"MedicalCare/pkg/e"
	"MedicalCare/pkg/logging"
)

type ArticleService struct {
	Title   string `json:"title" form:"title"`
	Desc    string `json:"desc" form:"desc"`
	Content string `json:"content" form:"content"`
}

// Create 新建文章
func (service *ArticleService) Create(uid uint) model.Response {
	code := e.Success

	var user model.User
	if err := model.DB.Model(model.User{}).Where("id = ?", uid).Find(&user).Error; err != nil {
		code = e.InvalidParams
		logging.Info(err)
		return model.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "此用户不存在",
		}
	}

	article := model.Article{
		User:    user,
		Title:   service.Title,
		Desc:    service.Desc,
		Content: service.Content,
	}

	if err := model.DB.Model(model.Article{}).Create(&article).Error; err != nil {
		code = e.Error
		logging.Info(err)
		return model.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "文章创建异常",
		}
	}

	return model.Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: "Successful write the article",
	}
}

// Remove 移出文章
// TODO Remove article api
func (service *ArticleService) Remove(bid, uid uint) model.Response {
	code := e.Success

	return model.Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: "Successful remove this article",
	}
}

// Update 更新文章信息
// TODO Update article api
func (service *ArticleService) Update(bid, uid uint) model.Response {
	code := e.Success

	return model.Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: "Successful update this article",
	}
}

// Show 展示指定文章
func (service *ArticleService) Show(bid uint) model.Response {
	code := e.Success

	var article model.Article
	if err := model.DB.Joins("User", model.DB.Where(&model.User{State: false})).Find(&article).Error; err != nil {
		code = e.InvalidParams
		logging.Info(err)
		return model.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "查找不到该文章",
		}
	}

	return model.Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: article,
	}
}
