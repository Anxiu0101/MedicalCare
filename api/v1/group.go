package v1

import (
	"MedicalCare/model"
	"MedicalCare/pkg/logging"
	"MedicalCare/pkg/util"
	"MedicalCare/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// CreateGroup godoc
// @Summary create group
// @Schemes
// @Description set up a group
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} CreateGroup
// @Router /group [post]
func CreateGroup(c *gin.Context) {
	var groupService service.GroupService
	if err := c.ShouldBind(&groupService); err == nil {
		res := groupService.Create()
		c.JSON(http.StatusOK, res)
	} else {
		logging.Info(err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(err))
	}
}

// InviteMember godoc
// @Summary invite member
// @Schemes
// @Description invite group members
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} InviteMember
// @Router /group/member [put]
func InviteMember(c *gin.Context) {
	var groupService service.GroupService
	if err := c.ShouldBind(&groupService); err == nil {
		claim, _ := util.ParseToken(c.GetHeader("Authorization"))
		res := groupService.Invite(claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		logging.Info(err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse(err))
	}
}
