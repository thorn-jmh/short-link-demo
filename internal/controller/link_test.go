package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/dao"
	"go-svc-tpl/internal/dao/ent"
	"go-svc-tpl/internal/dao/ent/link"
	"go-svc-tpl/utils/stacktrace"
	"net/http/httptest"
	"testing"
	"time"
)

var testLinkCtl = NewLinkController()

func mockData() {
	clearMock()
	ctx := context.TODO()
	// mock user
	dao.DB.User.Create().
		SetID(1).
		SetName("test").
		SetPassword(getHashedPassword("123456")).
		SetEmail("test@qwq.com").
		SaveX(ctx)
	dao.DB.User.Create().
		SetName("uuu").
		SetPassword(getHashedPassword("password")).
		SetEmail("uuu@test.com").
		SaveX(ctx)

	// mock link
	dao.DB.Link.Create().
		SetShort("simple").
		SetOrigin("https://www.google.com").
		SetComment("A short link for google").
		SetActive(true).
		SaveX(ctx)

	dao.DB.Link.Create().
		SetShort("withTime").
		SetOrigin("https://www.baidu.com").
		SetComment("A short link for baidu").
		SetActive(true).
		SetStartTime(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)).
		SetEndTime(time.Date(2021, 12, 31, 0, 0, 0, 0, time.Local)).
		SaveX(ctx)
	dao.DB.Link.Create().
		SetShort("withUser").
		SetOrigin("https://www.bing.com").
		SetComment("A short link for bing").
		SetActive(true).
		SetOwnerID(1).
		SaveX(ctx)
}

func clearMock() {
	ctx := context.TODO()
	dao.DB.User.Delete().ExecX(ctx)
	dao.DB.Link.Delete().ExecX(ctx)
}

func TestLinkController_Create(t *testing.T) {
	mockData()
	var tests = []struct {
		userID    uint
		input     dto.LinkCreateReq
		errorCode stacktrace.ErrorCode
	}{
		{
			input: dto.LinkCreateReq{
				Short:   "",
				Origin:  "https://www.baidu.com",
				Comment: "A short link for baidu",
			},
		},
		{
			input: dto.LinkCreateReq{
				Short:   "simple",
				Origin:  "https://www.baidu.com",
				Comment: "A short link for baidu",
			},
			errorCode: dto.ErrShortLinkExist,
		},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Set(USER_ID_KEY, test.userID)
		_, err := testLinkCtl.Create(ctx, &test.input)

		if test.errorCode != 0 {
			assert.Equal(t, test.errorCode, stacktrace.GetCode(err))
			continue
		}
		assert.NoError(t, err)
	}

	clearMock()
}

func TestLinkController_Delete(t *testing.T) {
	mockData()
	var tests = []struct {
		userID    uint
		input     dto.LinkDeleteReq
		errorCode stacktrace.ErrorCode
	}{
		{
			input: dto.LinkDeleteReq{
				Short: "withUser",
			},
			userID:    4,
			errorCode: dto.ErrPrivilege,
		},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Set(USER_ID_KEY, test.userID)
		err := testLinkCtl.Delete(ctx, &test.input)

		if test.errorCode != 0 {
			assert.Equal(t, test.errorCode, stacktrace.GetCode(err))
			continue
		}

		assert.NoError(t, err, "Delete link %q failed", test.input.Short)
		verifyDelete(t, ctx, test.input.Short)
	}
	clearMock()
}

func verifyDelete(t *testing.T, c context.Context, short string) {
	t.Helper()
	_, err := dao.DB.Link.Query().Where(link.Short(short)).Only(c)
	if assert.Errorf(t, err, "Link %q still exists after delete", short) {
		assert.True(t, ent.IsNotFound(err))
	}
}

func TestLinkController_GetInfo(t *testing.T) {
	mockData()
	var tests = []struct {
		input     dto.GetLinkInfoReq
		want      dto.GetLinkInfoResp
		errorCode stacktrace.ErrorCode
	}{
		{
			input: dto.GetLinkInfoReq{
				Short: "simple",
			},
			want: dto.GetLinkInfoResp{
				Short:   "simple",
				Origin:  "https://www.google.com",
				Comment: "A short link for google",
				Active:  true,
			},
		},
		{
			input: dto.GetLinkInfoReq{
				Short: "qwq",
			},
			errorCode: dto.ErrNoShortLink,
		},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		got, err := testLinkCtl.GetInfo(ctx, &test.input)

		if test.errorCode != 0 {
			assert.Equal(t, test.errorCode, stacktrace.GetCode(err))
			continue
		}
		assert.NoError(t, err)
		assert.EqualValues(t, &test.want, got)
	}
	clearMock()
}

func mockListData() {
	mockData()
	ctx := context.TODO()
	dao.DB.Link.Create().
		SetShort("test1").
		SetOrigin("https://www.google.com").
		SetOwnerID(1).
		SaveX(ctx)
	dao.DB.Link.Create().
		SetShort("test2").
		SetOrigin("https://www.google.com").
		SetOwnerID(1).
		SaveX(ctx)
	dao.DB.Link.Create().
		SetShort("test3").
		SetOrigin("https://www.google.com").
		SetOwnerID(1).
		SaveX(ctx)
	dao.DB.Link.Create().
		SetShort("test4").
		SetOrigin("https://www.google.com").
		SetOwnerID(1).
		SaveX(ctx)
	dao.DB.Link.Create().
		SetShort("test5").
		SetOrigin("https://www.google.com").
		SetOwnerID(1).
		SaveX(ctx)
	dao.DB.Link.Create().
		SetShort("test6").
		SetOrigin("https://www.google.com").
		SetOwnerID(1).
		SaveX(ctx)
	dao.DB.Link.Create().
		SetShort("test7").
		SetOrigin("https://www.google.com").
		SetOwnerID(1).
		SaveX(ctx)
	dao.DB.Link.Create().
		SetShort("test8").
		SetOrigin("https://www.google.com").
		SetOwnerID(1).
		SaveX(ctx)
	dao.DB.Link.Create().
		SetShort("test9").
		SetOrigin("https://www.google.com").
		SetOwnerID(1).
		SaveX(ctx)
	dao.DB.Link.Create().
		SetShort("test10").
		SetOrigin("https://www.google.com").
		SetOwnerID(1).
		SaveX(ctx)
}

func TestLinkController_List(t *testing.T) {
	mockListData()
	var tests = []struct {
		userID    uint
		input     dto.LinkListReq
		errorCode stacktrace.ErrorCode
	}{
		{
			input: dto.LinkListReq{
				PageNumber: 1,
				PageSize:   10,
			},
			errorCode: dto.ErrPrivilege,
		},
		{
			input: dto.LinkListReq{
				PageNumber: 1,
				PageSize:   -1,
			},
			userID: 1,
		},
		{
			input: dto.LinkListReq{
				PageNumber: 2,
				PageSize:   10,
			},
			userID: 1,
		},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Set(USER_ID_KEY, test.userID)
		_, err := testLinkCtl.List(ctx, &test.input)

		if test.errorCode != 0 {
			assert.Equal(t, test.errorCode, stacktrace.GetCode(err), "Unexpected error %v", err)
			continue
		}
		assert.NoError(t, err)
		// TODO: 嬾得寫了
	}
	clearMock()
}

func TestLinkController_UpdateInfo(t *testing.T) {
	mockData()
	var tests = []struct {
		userID    uint
		input     dto.UpdateLinkInfoReq
		errorCode stacktrace.ErrorCode
	}{
		{
			input: dto.UpdateLinkInfoReq{
				Short:   "withUser",
				Origin:  "https://www.qwqwq.com",
				Comment: "A short link for google",
				Active:  true,
			},
			userID: 1,
		},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Set(USER_ID_KEY, test.userID)
		err := testLinkCtl.UpdateInfo(ctx, &test.input)

		if test.errorCode != 0 {
			assert.Equal(t, test.errorCode, stacktrace.GetCode(err), "Update failed, input: %+v", test.input)
			continue
		}
		assert.NoError(t, err)
	}
	clearMock()
}
