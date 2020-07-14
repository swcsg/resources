package app_dependence

import (
	"resource/pkg/tools/http"
	"resource/pkg/types"
)

// AppDependenceController 有状态应用依赖资源接口封装
type AppDependenceController struct {
	HTTPTool      *http.HTTPTool
	AppDependence types.AppDependence
}
