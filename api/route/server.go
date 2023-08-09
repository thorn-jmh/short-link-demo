package route

import (
	"github.com/gin-gonic/gin"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/controller"
)

func setupServerController(r *gin.RouterGroup) {
	scr := ServerCtlWrapper{
		ctl: controller.NewServerController(),
	}

	r.GET("/:short", scr.ShortLinkServer)
	r.GET("/api/captcha/:target", scr.ctl.CaptchaServer)
}

type ServerCtlWrapper struct {
	ctl controller.IServerController
}

// Controller

func (w *ServerCtlWrapper) ShortLinkServer(c *gin.Context) {
	short := c.Param("short")
	err := w.ctl.ShortLinkServer(c, short)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
}
