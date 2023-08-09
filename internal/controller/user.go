package controller

import (
	"crypto/sha256"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/dao"
	"go-svc-tpl/internal/dao/ent"
	"go-svc-tpl/internal/dao/ent/user"
	"go-svc-tpl/utils/jwt"
	"go-svc-tpl/utils/stacktrace"
)

// >>>>>>>>>>>>>>>>>> Interface  >>>>>>>>>>>>>>>>>>

type IUserController interface {
	Register(*gin.Context, *dto.UserRegisterReq) error
	Captcha(*gin.Context) (*dto.UserCaptchaResp, error)
	Login(*gin.Context, *dto.UserLoginReq) error
	Logout(*gin.Context) error
	GetInfo(*gin.Context) (*dto.GetUserInfoResp, error)
	UpdateInfo(*gin.Context, *dto.UpdateUserInfoReq) error
	Passwd(*gin.Context, *dto.UserPasswdReq) error
}

// >>>>>>>>>>>>>>>>>> Controller >>>>>>>>>>>>>>>>>>

// check interface implementation
var _ IUserController = (*UserController)(nil)

var NewUserController = func() *UserController {
	var appCfg UserCtlCfg
	if err := viper.Sub("App").Unmarshal(&appCfg); err != nil {
		logrus.Fatal(err)
	}
	return &UserController{
		cfg: &appCfg,
	}
}

type UserCtlCfg struct {
	Origin       string `mapstructure:"Host"`
	CookieExpire int    `mapstructure:"CookieExpire"`
}

type UserController struct {
	cfg *UserCtlCfg
}

// ------------------ Register ------------------

func (c *UserController) Register(ctx *gin.Context, req *dto.UserRegisterReq) error {

	_, err := dao.DB.User.Create().
		SetName(req.Name).
		SetEmail(req.Email).
		SetPassword(getHashedPassword(req.Password)).
		Save(ctx)

	if ent.IsConstraintError(err) {
		return stacktrace.PropagateWithCode(err, dto.ErrEmailExist, "Email has been registered")
	}

	return err
}

func getHashedPassword(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}

// ------------------ Captcha ------------------

func (c *UserController) Captcha(ctx *gin.Context) (*dto.UserCaptchaResp, error) {
	cid := captcha.New()
	return &dto.UserCaptchaResp{
		CaptchaID:  cid,
		CaptchaURL: c.cfg.Origin + "/api/captcha/" + cid + ".png",
	}, nil
}

// ------------------ Login ------------------

func (c *UserController) Login(ctx *gin.Context, req *dto.UserLoginReq) error {

	if !captcha.VerifyString(req.CaptchaID, req.CaptchaValue) {
		return stacktrace.NewErrorWithCode(dto.ErrCaptcha, "Wrong captcha")
	}

	u, err := dao.DB.User.Query().
		Where(user.Email(req.Email)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return stacktrace.PropagateWithCode(err, dto.ErrUserNotFound, "User not found")
		}
	}

	if u.Password != getHashedPassword(req.Password) {
		return stacktrace.NewErrorWithCode(dto.ErrPassword, "Wrong password")
	}

	token := jwt.Release(&jwt.Claims{UserID: u.ID})
	ctx.SetCookie(COOKIE_NAME, token, c.cfg.CookieExpire, "/", c.cfg.Origin, false, true)
	return nil
}

// ------------------ Logout ------------------

func (c *UserController) Logout(ctx *gin.Context) error {
	ctx.SetCookie(COOKIE_NAME, "", -1, "/", c.cfg.Origin, false, true)
	return nil
}

// ------------------ GetInfo ------------------

func (c *UserController) GetInfo(ctx *gin.Context) (*dto.GetUserInfoResp, error) {
	userID := ctx.GetUint(USER_ID_KEY)
	u, err := dao.DB.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &dto.GetUserInfoResp{
		Name:  u.Name,
		Email: u.Email,
		ID:    u.ID,
	}, nil
}

// ------------------ UpdateInfo ------------------

func (c *UserController) UpdateInfo(ctx *gin.Context, req *dto.UpdateUserInfoReq) error {
	userID := ctx.GetUint(USER_ID_KEY)
	_, err := dao.DB.User.Update().
		Where(user.ID(userID)).
		SetName(req.Name).
		SetEmail(req.Email).
		Save(ctx)

	if ent.IsConstraintError(err) {
		return stacktrace.PropagateWithCode(err, dto.ErrEmailExist, "Email has been registered")
	}
	return err
}

// ------------------ Passwd ------------------

func (c *UserController) Passwd(ctx *gin.Context, req *dto.UserPasswdReq) error {
	userID := ctx.GetUint(USER_ID_KEY)
	u, err := dao.DB.User.Get(ctx, userID)
	if err != nil {
		return err
	}

	if u.Password != getHashedPassword(req.OldPwd) {
		return stacktrace.NewErrorWithCode(dto.ErrPassword, "Wrong password")
	}

	_, err = dao.DB.User.Update().
		Where(user.ID(userID)).
		SetPassword(getHashedPassword(req.NewPwd)).
		Save(ctx)

	return err
}
