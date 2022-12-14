// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-test/internal/dao/internal"
)

// internalPrAuthK8SAuthDao is internal type for wrapping internal DAO implements.
type internalPrAuthK8SAuthDao = *internal.PrAuthK8SAuthDao

// prAuthK8SAuthDao is the data access object for table tt_pr_auth_k8s_auth.
// You can define custom methods on it to extend its functionality as you wish.
type prAuthK8SAuthDao struct {
	internalPrAuthK8SAuthDao
}

var (
	// PrAuthK8SAuth is globally public accessible object for table tt_pr_auth_k8s_auth operations.
	PrAuthK8SAuth = prAuthK8SAuthDao{
		internal.NewPrAuthK8SAuthDao(),
	}
)

// Fill with you ideas below.
