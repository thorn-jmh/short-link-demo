package route

import (
	"github.com/gin-gonic/gin"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/controller"
)

func setupUserController(r *gin.RouterGroup) {
	ucw := UserCtlWrapper{
		ctl: controller.NewUserController(),
	}

	p := r.Group("/user")
	p.POST("/register", ucw.Register)
	p.GET("/captcha", ucw.Captcha)
	p.POST("/login", ucw.Login)
	p.POST("/logout", controller.BasicAuthMidware(), ucw.Logout)
	p.GET("/info", controller.BasicAuthMidware(), ucw.GetInfo)
	p.POST("/info", controller.BasicAuthMidware(), ucw.UpdateInfo)
	p.POST("/passwd", controller.BasicAuthMidware(), ucw.Passwd)
}

type UserCtlWrapper struct {
	ctl controller.IUserController
}

// >>>>>>>>>>>>>>>>>> Controller >>>>>>>>>>>>>>>>>>

func (w *UserCtlWrapper) Register(ctx *gin.Context) {
	var req dto.UserRegisterReq
	if err := dto.BindReq(ctx, &req); err != nil {
		dto.ResponseFail(ctx, err)
		return
	}
	err := w.ctl.Register(ctx, &req)
	if err != nil {
		dto.ResponseFail(ctx, err)
		return
	}
	dto.ResponseSuccess(ctx, nil)
}

func (w *UserCtlWrapper) Captcha(ctx *gin.Context) {
	resp, err := w.ctl.Captcha(ctx)
	if err != nil {
		dto.ResponseFail(ctx, err)
		return
	}
	dto.ResponseSuccess(ctx, resp)
}

func (w *UserCtlWrapper) Login(ctx *gin.Context) {
	var req dto.UserLoginReq
	if err := dto.BindReq(ctx, &req); err != nil {
		dto.ResponseFail(ctx, err)
		return
	}
	err := w.ctl.Login(ctx, &req)
	if err != nil {
		dto.ResponseFail(ctx, err)
		return
	}
	dto.ResponseSuccess(ctx, nil)
}

func (w *UserCtlWrapper) Logout(ctx *gin.Context) {
	err := w.ctl.Logout(ctx)
	if err != nil {
		dto.ResponseFail(ctx, err)
		return
	}
	dto.ResponseSuccess(ctx, nil)
}

func (w *UserCtlWrapper) GetInfo(ctx *gin.Context) {
	resp, err := w.ctl.GetInfo(ctx)
	if err != nil {
		dto.ResponseFail(ctx, err)
		return
	}
	dto.ResponseSuccess(ctx, resp)
}

func (w *UserCtlWrapper) UpdateInfo(ctx *gin.Context) {
	var req dto.UpdateUserInfoReq
	if err := dto.BindReq(ctx, &req); err != nil {
		dto.ResponseFail(ctx, err)
		return
	}
	err := w.ctl.UpdateInfo(ctx, &req)
	if err != nil {
		dto.ResponseFail(ctx, err)
		return
	}
	dto.ResponseSuccess(ctx, nil)
}

func (w *UserCtlWrapper) Passwd(ctx *gin.Context) {
	var req dto.UserPasswdReq
	if err := dto.BindReq(ctx, &req); err != nil {
		dto.ResponseFail(ctx, err)
		return
	}
	err := w.ctl.Passwd(ctx, &req)
	if err != nil {
		dto.ResponseFail(ctx, err)
		return
	}
	dto.ResponseSuccess(ctx, nil)
}
