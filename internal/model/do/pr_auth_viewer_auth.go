// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PrAuthViewerAuth is the golang structure of table tt_pr_auth_viewer_auth for DAO operations like Where/Data.
type PrAuthViewerAuth struct {
	g.Meta    `orm:"table:tt_pr_auth_viewer_auth, do:true"`
	Id        interface{} //
	AuthName  interface{} //
	Authority interface{} //
	Desc      interface{} //
	Kind      interface{} //
	Rank      interface{} //
}
