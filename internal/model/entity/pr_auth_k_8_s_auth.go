// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// PrAuthK8SAuth is the golang structure for table pr_auth_k8s_auth.
type PrAuthK8SAuth struct {
	Id           string `json:"id"           description:""`
	AuthName     string `json:"authName"     description:"权限名, 如:svc-r, deploy-{name}-r"`
	Authority    string `json:"authority"    description:"R | W | X"`
	Kind         string `json:"kind"         description:"资源种类"`
	ClusterId    int    `json:"clusterId"    description:""`
	ClusterName  string `json:"clusterName"  description:""`
	Namespace    string `json:"namespace"    description:""`
	Desc         string `json:"desc"         description:"描述"`
	ResourceName string `json:"resourceName" description:"Resource Name"`
}
