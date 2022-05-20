package v1

import (
	"MedicalCare/cache"
	"MedicalCare/pkg/util"
	"MedicalCare/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// Chat godoc
// @Summary      聊天室接口
// @Description  用于聊天室的使用
// @Tags         聊天室相关接口
// @Accept       json
// @Produce      json
// @Security     x-token
// @Param        x-token   header    string  true  "Authorization"
// @Param        receiver  path      string  true  "聊天对象"
// @Success      200       {object}  model.Response
// @Failure      400       {object}  model.Response
// @Failure      500       {object}  model.Response
// @Router       /chat/:receiver [get]
func Chat(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	sender := claim.Username
	receiver := c.Param("receiver")

	// 将 http 协议升级为 websocket 协议
	conn, err := (&websocket.Upgrader{
		// CheckOrigin 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
			return true
		}}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	log.Println("webSocket 建立连接:", conn.RemoteAddr().String())

	// Create A new Client Object
	client := &service.Client{
		SID:    sender,
		RID:    receiver,
		Socket: conn,
		Send:   make(chan []byte),
	}

	// 登入时将用户加载入内存
	cache.RedisClient.Incr(cache.Ctx, client.SID)

	// 用户注册到用户管理上
	service.Manager.Register <- client
	go client.Read()
	go client.Write()
}
