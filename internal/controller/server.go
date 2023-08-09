package controller

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/dao"
	"go-svc-tpl/internal/dao/ent"
	"go-svc-tpl/internal/dao/ent/link"
	"go-svc-tpl/utils/stacktrace"
	"net/http"
	"time"
)

type IServerController interface {
	ShortLinkServer(*gin.Context, string) error
	CaptchaServer(ctx *gin.Context)
}

// >>>>>>>>>>>>>>>>>> Controller >>>>>>>>>>>>>>>>>>

// check interface implementation
var _ IServerController = (*ServerController)(nil)

var NewServerController = func() *ServerController {
	var appCfg ServerCtlCfg
	if err := viper.Sub("App").Unmarshal(&appCfg); err != nil {
		logrus.Fatal(err)
	}
	return &ServerController{
		captchaServer: captcha.Server(appCfg.CaptchaWidth, appCfg.CaptchaHeight),
		cfg:           &appCfg,
	}
}

type ServerCtlCfg struct {
	CaptchaWidth  int `mapstructure:"CaptchaWidth"`
	CaptchaHeight int `mapstructure:"CaptchaHeight"`
}

type ServerController struct {
	// maybe some logic config to read from viper
	// or a service dependency
	captchaServer http.Handler
	cfg           *ServerCtlCfg
}

func (c *ServerController) ShortLinkServer(ctx *gin.Context, short string) error {
	l, err := dao.DB.Link.Query().Where(link.Short(short)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return stacktrace.PropagateWithCode(err, dto.ErrNoShortLink, "Short link not found")
		}
		return err
	}

	if !l.Active {
		return stacktrace.PropagateWithCode(err, dto.ErrShortLinkActive, "Short link not active")
	}

	if l.StartTime != nil && l.StartTime.After(time.Now()) {
		return stacktrace.PropagateWithCode(err, dto.ErrShortLinkTime, "Short link not start")
	}
	if l.EndTime != nil && l.EndTime.Before(time.Now()) {
		return stacktrace.PropagateWithCode(err, dto.ErrShortLinkTime, "Short link expired")
	}

	ctx.Redirect(307, l.Origin)
	return nil
}

func (c *ServerController) CaptchaServer(ctx *gin.Context) {
	c.captchaServer.ServeHTTP(ctx.Writer, ctx.Request)
}
