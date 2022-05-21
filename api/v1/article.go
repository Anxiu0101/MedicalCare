package v1

import (
	"MedicalCare/model"
	"MedicalCare/pkg/logging"
	"MedicalCare/pkg/util"
	"MedicalCare/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// WriteArticle godoc
// @Summary      文章发布接口
// @Description  用于新建文章
// @Tags         文章相关接口
// @Accept       json
// @Produce      json
// @Security     x-token
// @Param        x-token  header    string  true   "Authorization"
// @Param        title    formData  string  true   "文章标题"
// @Param        desc     formData  string  false  "文章简介"
// @Param        content  formData  string  true   "文章内容"
// @Success      200      {object}  model.Response
// @Failure      400      {object}  model.Response
// @Failure      500      {object}  model.Response
// @Router       /api/v1/blog [post]
func WriteArticle(c *gin.Context) {
	var articleService service.ArticleService
	if err := c.ShouldBind(&articleService); err == nil {
		claim, _ := util.ParseToken(c.GetHeader("Authorization"))
		res := articleService.Create(claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		logging.Info(err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(err))
	}
}

func RemoveArticle(c *gin.Context) {

}

func UpdateArticle(c *gin.Context) {

}

// GetArticle godoc
// @Summary      文章阅读接口
// @Description  用于获取指定文章
// @Tags         文章相关接口
// @Accept       json
// @Produce      json
// @Security     x-token
// @Param        x-token  header    string  true   "Authorization"
// @Param        bid    path  string  true   "指定文章的ID"
// @Success      200      {object}  model.Response{data=model.Article{User=model.User}}
// @Failure      400      {object}  model.Response
// @Failure      500      {object}  model.Response
// @Router       /api/v1/blog [post]
func GetArticle(c *gin.Context) {
	var articleService service.ArticleService
	if err := c.ShouldBind(&articleService); err == nil {
		bid, _ := strconv.Atoi(c.Param("bid"))
		res := articleService.Show(uint(bid))
		c.JSON(http.StatusOK, res)
	} else {
		logging.Info(err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(err))
	}
}

func GetArticles(c *gin.Context) {

}
