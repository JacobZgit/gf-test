// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OpRecord is the golang structure for table op_record.
type OpRecord struct {
	Id          uint        `json:"id"          description:""`
	Username    string      `json:"username"    description:""`
	Uri         string      `json:"uri"         description:""`
	Method      string      `json:"method"      description:""`
	Code        int         `json:"code"        description:""`
	Header      string      `json:"header"      description:""`
	QueryParams string      `json:"queryParams" description:""`
	QueryBody   string      `json:"queryBody"   description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:""`
	DeletedAt   *gtime.Time `json:"deletedAt"   description:""`
}