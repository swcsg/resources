package crd_dependence

import (
	"resource/pkg/tools/http"
	"resource/pkg/types"
)

type CRDDependenceController struct {
	HttpTool      *http.HttpTool
	CRDDependence types.CRDDependence
}
