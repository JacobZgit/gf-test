// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PrAuthRole is the golang structure for table pr_auth_role.
type PrAuthRole struct {
	Id        uint        `json:"id"        description:""`
	RoleName  string      `json:"roleName"  description:"role name"`
	Desc      string      `json:"desc"      description:""`
	Kind      string      `json:"kind"      description:"权限种类(viewer, k8s)"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
