package controller

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/dao"
	"go-svc-tpl/internal/dao/ent"
	"go-svc-tpl/internal/dao/ent/link"
	"go-svc-tpl/internal/dao/ent/user"
	"go-svc-tpl/utils/stacktrace"
	"gopkg.in/guregu/null.v4"
)

// >>>>>>>>>>>>>>>>>> Interface  >>>>>>>>>>>>>>>>>>

type ILinkController interface {
	Create(*gin.Context, *dto.LinkCreateReq) (*dto.LinkCreateResp, error)
	Delete(*gin.Context, *dto.LinkDeleteReq) error
	GetInfo(*gin.Context, *dto.GetLinkInfoReq) (*dto.GetLinkInfoResp, error)
	UpdateInfo(*gin.Context, *dto.UpdateLinkInfoReq) error
	List(*gin.Context, *dto.LinkListReq) (*dto.LinkListResp, error)
}

// >>>>>>>>>>>>>>>>>> Controller >>>>>>>>>>>>>>>>>>

// check interface implementation
var _ ILinkController = (*LinkController)(nil)

var NewLinkController = func() *LinkController {
	return &LinkController{}
}

type LinkController struct {
	// maybe some logic config to read from viper
	// or a service dependency
}

// ---------------------- Create ----------------------

func (c *LinkController) Create(ctx *gin.Context, req *dto.LinkCreateReq) (*dto.LinkCreateResp, error) {
	userID := ctx.GetUint(USER_ID_KEY)
	var nillableUserID *uint = nil
	if userID != 0 {
		nillableUserID = &userID
	}

	l, err := dao.DB.Link.Create().
		SetComment(req.Comment).
		SetOrigin(req.Origin).
		SetNillableStartTime(req.StartTime.Ptr()).
		SetNillableEndTime(req.EndTime.Ptr()).
		SetActive(true).
		SetNillableOwnerID(nillableUserID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	l.Short = req.Short
	if l.Short == "" {
		l.Short = generateShort(l.ID)
	}

	_, err = l.Update().
		SetShort(l.Short).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, stacktrace.PropagateWithCode(err, dto.ErrShortLinkExist, "Short link already exists")
		}
		return nil, err
	}

	return &dto.LinkCreateResp{
		Short:     l.Short,
		Origin:    l.Origin,
		Comment:   l.Comment,
		StartTime: null.TimeFromPtr(l.StartTime),
		EndTime:   null.TimeFromPtr(l.EndTime),
		Active:    l.Active,
	}, nil
}

// base62 char table
var charTable = []byte("Ace7BDFGHIJKL2MNOPQRTUVWXYZabcdfghijkmnopqrstuvwxyz015689")

func _10to62(num uint) []byte {
	var result []byte
	for num > 0 {
		result = append(result, charTable[num%62])
		num /= 62
	}
	return result
}

func generateShort(id uint) string {
	hash := _10to62(id)
	return string(hash)
}

// ---------------------- Delete ----------------------

func (c *LinkController) Delete(ctx *gin.Context, req *dto.LinkDeleteReq) error {
	userID := ctx.GetUint(USER_ID_KEY)

	l, err := dao.DB.Link.Query().
		Where(link.HasOwnerWith(user.ID(userID))).
		Where(link.Short(req.Short)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return stacktrace.PropagateWithCode(err, dto.ErrPrivilege, "This link does not belong to you")
		}
		return err
	}

	err = dao.DB.Link.DeleteOne(l).Exec(ctx)

	return err
}

// ---------------------- GetInfo ----------------------

func (c *LinkController) GetInfo(ctx *gin.Context, req *dto.GetLinkInfoReq) (*dto.GetLinkInfoResp, error) {
	l, err := dao.DB.Link.Query().
		Where(link.Short(req.Short)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, stacktrace.PropagateWithCode(err, dto.ErrNoShortLink, "Link does not exist")
		}
		return nil, err
	}

	return &dto.GetLinkInfoResp{
		Short:     l.Short,
		Origin:    l.Origin,
		Comment:   l.Comment,
		StartTime: null.TimeFromPtr(l.StartTime),
		EndTime:   null.TimeFromPtr(l.EndTime),
		Active:    l.Active,
	}, nil
}

// ---------------------- UpdateInfo ----------------------

func (c *LinkController) UpdateInfo(ctx *gin.Context, req *dto.UpdateLinkInfoReq) error {
	userID := ctx.GetUint(USER_ID_KEY)

	l, err := dao.DB.Link.Query().
		Where(link.HasOwnerWith(user.ID(userID))).
		Where(link.Short(req.Short)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return stacktrace.PropagateWithCode(err, dto.ErrPrivilege, "This link does not belong to you")
		}
		return err
	}

	_, err = l.Update().
		SetShort(req.Short).
		SetOrigin(req.Origin).
		SetComment(req.Comment).
		SetNillableStartTime(req.StartTime.Ptr()).
		SetNillableEndTime(req.EndTime.Ptr()).
		SetActive(req.Active).
		Save(ctx)

	return err
}

// ---------------------- List ----------------------

func (c *LinkController) List(ctx *gin.Context, req *dto.LinkListReq) (*dto.LinkListResp, error) {
	userID := ctx.GetUint(USER_ID_KEY)
	if userID == 0 {
		return nil, stacktrace.NewErrorWithCode(dto.ErrPrivilege, "You are not logged in")
	}

	count, err := dao.DB.User.Query().
		Where(user.ID(userID)).
		QueryLinks().
		Count(ctx)

	if err != nil {
		return nil, err
	}

	offset, limit := calculatePage(req.PageNumber, req.PageSize, count)

	ls, err := dao.DB.User.Query().
		Where(user.ID(userID)).
		QueryLinks().
		Order(link.ByID()).
		Offset(offset).
		Limit(limit).
		All(ctx)

	var resp dto.LinkListResp
	resp.Total = count

	linq.From(ls).SelectT(func(l *ent.Link) dto.ShortLinkModel {
		return dto.ShortLinkModel{
			Short:     l.Short,
			Origin:    l.Origin,
			Comment:   l.Comment,
			StartTime: null.TimeFromPtr(l.StartTime),
			EndTime:   null.TimeFromPtr(l.EndTime),
			Active:    l.Active,
		}
	}).ToSlice(&resp.Links)

	return &resp, nil
}

func calculatePage(pageNumber, pageSize, count int) (int, int) {
	offset := (pageNumber - 1) * pageSize
	limit := pageSize
	if offset > count {
		logrus.Warnf("offset %d is larger than count %d", offset, count)
	}

	if limit < 0 {
		limit = count
	}
	if offset < 0 {
		offset = 0
	}
	return offset, limit
}
