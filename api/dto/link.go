package dto

import (
	"gopkg.in/guregu/null.v4"
)

// >>>>>>>>>> LinkCreate >>>>>>>>>>

type LinkCreateReq struct {
	Comment   string    `json:"comment"`    // 备注信息
	EndTime   null.Time `json:"end_time"`   // 到期时间，UTC
	Origin    string    `json:"origin"`     // 原始链接
	Short     string    `json:"short"`      // 短链ID，全局唯一
	StartTime null.Time `json:"start_time"` // 起始时间，UTC
}

type LinkCreateResp struct {
	Active    bool      `json:"active"`     // 服务状态
	Comment   string    `json:"comment"`    // 备注信息
	EndTime   null.Time `json:"end_time"`   // 到期时间，UTC
	Origin    string    `json:"origin"`     // 原始链接
	Short     string    `json:"short"`      // 短链ID，全局唯一
	StartTime null.Time `json:"start_time"` // 起始时间，UTC
}

// >>>>>>>>>> LinkDelete >>>>>>>>>>

type LinkDeleteReq struct {
	Short string `form:"short"` // 短链接的唯一 ID
}

// >>>>>>>>>> GetLinkInfo >>>>>>>>>>

type GetLinkInfoReq struct {
	Short string `form:"short"` // 短链接的唯一 ID
}

type GetLinkInfoResp struct {
	Active    bool      `json:"active"`     // 服务状态
	Comment   string    `json:"comment"`    // 备注信息
	EndTime   null.Time `json:"end_time"`   // 到期时间，UTC
	Origin    string    `json:"origin"`     // 原始链接
	Short     string    `json:"short"`      // 短链ID，全局唯一
	StartTime null.Time `json:"start_time"` // 起始时间，UTC
}

// >>>>>>>>>> UpdateLinkInfo >>>>>>>>>>

type UpdateLinkInfoReq struct {
	Active    bool      `json:"active"`     // 服务状态
	Comment   string    `json:"comment"`    // 备注信息
	EndTime   null.Time `json:"end_time"`   // 到期时间，UTC
	Origin    string    `json:"origin"`     // 原始链接
	Short     string    `json:"short"`      // 短链ID，全局唯一
	StartTime null.Time `json:"start_time"` // 起始时间，UTC
}

// >>>>>>>>>> LinkList >>>>>>>>>>

type LinkListReq struct {
	PageNumber int `form:"page_number"` // 请求页码
	PageSize   int `form:"page_size"`   // 每页 item 数，-1 时代表全部返回
}

type LinkListResp struct {
	Links []ShortLinkModel `json:"links"`
	Total int              `json:"total"` // 总 item 数量
}

type ShortLinkModel struct {
	Active    bool      `json:"active"`     // 服务状态
	Comment   string    `json:"comment"`    // 备注信息
	EndTime   null.Time `json:"end_time"`   // 到期时间，UTC
	Origin    string    `json:"origin"`     // 原始链接
	Short     string    `json:"short"`      // 短链ID，全局唯一
	StartTime null.Time `json:"start_time"` // 起始时间，UTC
}
