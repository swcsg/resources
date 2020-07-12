package app_dependence

import (
	"resource/pkg/tools/http"
	"resource/pkg/types"
)

type AppDependenceController struct {
	HttpTool      *http.HttpTool
	AppDependence types.AppDependence
}
