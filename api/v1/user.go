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

// @BasePath  /api/v1

// UserRegister godoc
// @Summary      用户注册接口
// @Description  用于注册新用户，传入 username 和 password 作为新用户的 username 和 password
// @Tags         用户相关接口
// @Accept       application/json
// @Produce      application/json
// @Param        username  formData  string  true  "用户名"
// @Param        password  formData  string  true  "用户密码"
// @Success      200       {object}  model.Response{code: 200, data: "Successful Register", msg: "ok"}
// @Failure      400       "输入参数错误"
// @Failure      500       "服务错误"
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
// @Accept       application/json
// @Produce      application/json
// @Param        username  formData  string  true  "用户名"
// @Param        password  formData  string  true  "用户密码"
// @Success      200       {object}  model.Response{data:model.TokenData}
// @Failure      400       "输入参数错误"
// @Failure      500       "服务错误"
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
// @Accept       application/json
// @Produce      application/json
// @Param        password  formData  string  true  "新密码"
// @Param        Authorization  header  string  true  "token"
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
// @Summary  Get User Info
// @Schemes
// @Description  get user account info
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {string}  GetUserInfo
// @Router       /user/info [get]
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
// @Summary  Update User Info
// @Schemes
// @Description  update user account info
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {string}  UpdateUserInfo
// @Router       /user/info [post]
func UpdateUserInfo(c *gin.Context) {
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

// UserOnline godoc
// @Summary  Get User online
// @Schemes
// @Description  get user is online or not
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {string}  UserOnline
// @Router       /user/online [get]
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
