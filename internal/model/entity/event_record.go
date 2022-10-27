// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventRecord is the golang structure for table event_record.
type EventRecord struct {
	Id          uint        `json:"id"          description:""`
	Subject     string      `json:"subject"     description:"标题"`
	Version     string      `json:"version"     description:"版本"`
	Eid         string      `json:"eid"         description:"event id"`
	Source      string      `json:"source"      description:"事件源"`
	Type        string      `json:"type"        description:"事件类型"`
	Time        *gtime.Time `json:"time"        description:"触发时间"`
	ContentType string      `json:"contentType" description:"内容类型"`
	Data        string      `json:"data"        description:"数据参数"`
	IsRespond   int         `json:"isRespond"   description:"是否有响应函数"`
	Resp        string      `json:"resp"        description:"响应数据(暂时不记录)"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:""`
	DeletedAt   *gtime.Time `json:"deletedAt"   description:""`
}
