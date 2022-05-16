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

// @BasePath /api/v1

// UserRegister godoc
// @Summary User Register
// @Schemes
// @Description help User Create Account
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} UserRegister
// @Router /user/register [post]
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
// @Summary User Login
// @Schemes
// @Description help User login site
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} UserLogin
// @Router /user/login [post]
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
// @Summary User reset password
// @Schemes
// @Description help User update account password
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} ResetUserPassword
// @Router /user/password [post]
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
// @Summary Get User Info
// @Schemes
// @Description get user account info
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} GetUserInfo
// @Router /user/info [get]
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
// @Summary Update User Info
// @Schemes
// @Description update user account info
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} UpdateUserInfo
// @Router /user/info [post]
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
// @Summary Get User online
// @Schemes
// @Description get user is online or not
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} UserOnline
// @Router /user/online [get]
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
