package v1

import (
	"MedicalCare/cache"
	"MedicalCare/model"
	"MedicalCare/pkg/logging"
	"MedicalCare/pkg/util"
	"MedicalCare/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegister godoc
// @Summary      用户注册接口
// @Description  用于注册新用户，传入 username 和 password 作为新用户的 username 和 password
// @Tags         用户相关接口
// @Accept       json
// @Produce      json
// @Param        username  formData  string  true  "用户名"
// @Param        password  formData  string  true  "用户密码"
// @Success      200       {object}  model.Response
// @Failure      400       {object}  model.Response
// @Failure      500       {object}  model.Response
// @Router       /api/v1/user/register [post]
func UserRegister(c *gin.Context) {
	var accountService service.AccountService
	if err := c.ShouldBind(&accountService); err == nil {
		res := accountService.Register()
		c.JSON(http.StatusOK, res)
	} else {
		logging.Info(err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(err))
	}
}

// UserLogin godoc
// @Summary      用户登陆接口
// @Description  用于用户登陆，接受用户名和用户密码，将 access token 和 refresh token 作为结果返回，同时返回用户的账户信息
// @Tags         用户相关接口
// @Accept       json
// @Produce      json
// @Param        username  formData  string  true  "用户名"
// @Param        password  formData  string  true  "用户密码"
// @Success      200       {object}  model.Response{data=model.TokenData{User=model.AccountInfo}}
// @Failure      400       {object}  model.Response
// @Failure      500       {object}  model.Response
// @Router       /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	var accountService service.AccountService
	if err := c.ShouldBind(&accountService); err == nil {
		res := accountService.Login()
		c.JSON(http.StatusOK, res)
	} else {
		logging.Info(err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(err))
	}
}

// ResetUserPassword godoc
// @Summary      用户重新设置密码接口
// @Description  用于用户更新密码，需要新密码做为传入的参数，需要 token 验证
// @Tags         用户相关接口
// @Accept       json
// @Produce      json
// @Security     x-token
// @Param        x-token   header    string  true  "Authorization"
// @Param        password  formData  string  true  "新密码"
// @Success      200       {object}  model.Response
// @Failure      400       {object}  model.Response
// @Failure      500       {object}  model.Response
// @Router       /api/v1/user/password [post]
func ResetUserPassword(c *gin.Context) {
	var accountService service.AccountService
	if err := c.ShouldBind(&accountService); err == nil {
		claim, _ := util.ParseToken(c.GetHeader("Authorization"))
		res := accountService.ResetPassword(claim.ID, accountService.Password)
		c.JSON(http.StatusOK, res)
	} else {
		logging.Info(err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(err))
	}
}

func RefreshAccessToken(c *gin.Context) {

}

// GetUserInfo godoc
// @Summary      用户资料获取接口
// @Description  用于用户获取其个人资料
// @Tags         用户相关接口
// @Accept       json
// @Produce      json
// @Security     x-token
// @Param        x-token  header    string  true  "Authorization"
// @Success      200      {object}  model.Response{data=service.UserService}
// @Failure      400      {object}  model.Response
// @Failure      500      {object}  model.Response
// @Router       /api/v1/user/info [get]
func GetUserInfo(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err == nil {
		claim, _ := util.ParseToken(c.GetHeader("Authorization"))
		res := userService.GetInfo(claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		logging.Info(err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(err))
	}
}

// UpdateUserInfo godoc
// @Summary      用户资料更新接口
// @Description  用于用户更新其个人资料
// @Tags         用户相关接口
// @Accept       json
// @Produce      json
// @Security     x-token
// @Param        x-token  header    string  true   "Authorization"
// @Param        email    formData  string  false  "用户邮箱"
// @Param        gender   formData  int     false  "用户性别"
// @Param        age      formData  int     false  "用户年龄"
// @Param        tel      formData  string  false  "联系方式"
// @Success      200      {object}  model.Response
// @Failure      400      {object}  model.Response
// @Failure      500      {object}  model.Response
// @Router       /api/v1/user/info [put]
func UpdateUserInfo(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err == nil {
		claim, _ := util.ParseToken(c.GetHeader("Authorization"))
		res := userService.UpdateInfo(claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		logging.Info(err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(err))
	}
}

func UserOnline(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if cache.GetUserOnline(int(claim.ID)) {
		c.JSON(http.StatusOK, gin.H{
			"Data": "user is online",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Data": "user is offline",
		})
	}
}
