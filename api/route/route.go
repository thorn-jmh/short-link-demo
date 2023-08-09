package route

import (
	"github.com/gin-gonic/gin"
	"go-svc-tpl/api/dto"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Resp{
		Code: http.StatusOK,
		Msg:  "success",
		Data: "pong~",
	})
}

func SetupRouter(r *gin.RouterGroup) {
	api := r.Group("/api")
	api.GET("/ping", Ping)
	setupServerController(r)
	setupLinkController(api)
	setupUserController(api)
}
