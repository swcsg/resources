package crd_dependence

import (
	"resource/pkg/tools/http"
	"resource/pkg/types"
)

// CRDDependenceController 无状态应用依赖资源接口封装
type CRDDependenceController struct {
	HTTPTool      *http.HTTPTool
	CRDDependence types.CRDDependence
}
