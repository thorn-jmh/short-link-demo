package route

import (
	"github.com/gin-gonic/gin"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/controller"
)

func setupLinkController(r *gin.RouterGroup) {
	lcw := LinkCtlWrapper{
		ctl: controller.NewLinkController(),
	}

	p := r.Group("/link")
	p.POST("/create", controller.ParseTokenMidware(), lcw.Create)
	p.POST("/delete", controller.BasicAuthMidware(), lcw.Delete)
	p.GET("/info", lcw.GetInfo)
	p.POST("/info", controller.BasicAuthMidware(), lcw.UpdateInfo)
	p.GET("/list", controller.BasicAuthMidware(), lcw.List)
}

type LinkCtlWrapper struct {
	ctl controller.ILinkController
}

// >>>>>>>>>> Controller >>>>>>>>>>

func (w *LinkCtlWrapper) Create(c *gin.Context) {
	var req dto.LinkCreateReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.Create(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

func (w *LinkCtlWrapper) Delete(c *gin.Context) {
	var req dto.LinkDeleteReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	err := w.ctl.Delete(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, nil)
}

func (w *LinkCtlWrapper) GetInfo(c *gin.Context) {
	var req dto.GetLinkInfoReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.GetInfo(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

func (w *LinkCtlWrapper) UpdateInfo(c *gin.Context) {
	var req dto.UpdateLinkInfoReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	err := w.ctl.UpdateInfo(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, nil)
}

func (w *LinkCtlWrapper) List(c *gin.Context) {
	var req dto.LinkListReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.List(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}
