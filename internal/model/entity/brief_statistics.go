// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BriefStatistics is the golang structure for table brief_statistics.
type BriefStatistics struct {
	Id           int         `json:"id"           description:"主键"`
	Time         *gtime.Time `json:"time"         description:"日期"`
	PodCount     int         `json:"podCount"     description:"当日的pod总数"`
	NodeCount    int         `json:"nodeCount"    description:""`
	ClusterCount int         `json:"clusterCount" description:""`
	ServiceCount int         `json:"serviceCount" description:""`
}
