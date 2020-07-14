package crd_dependence

import (
	"resource/pkg/service/crd_dependence"
	"resource/pkg/tools/kubernetes"

	"github.com/gin-gonic/gin"
)

// CRDDependenceOptions 无状态应用依赖资源请求体
type CRDDependenceOptions struct {
	DeploymentName string `json:"deploymentName"`
	Namespace      string `json:"namespace"`
}

// CRDDependence 无状态应用依赖资源控制层
func CRDDependence(c *gin.Context) {

	var options = &CRDDependenceOptions{}

	var dc = &CRDDependenceController{
		CRDDependence: &crd_dependence.CRDDependence{
			StatefulSetClient:  &kubernetes.StatefulSet{},
			PodClient:          &kubernetes.PodClinet{},
			ServiceClient:      &kubernetes.ServiceClient{},
			PvcClient:          &kubernetes.PvcClient{},
			PvClient:           &kubernetes.PvClient{},
			SecretClient:       &kubernetes.SecretClient{},
			StorageClassClient: &kubernetes.StroageClassClient{},
			ConfigMapClient:    &kubernetes.ConfigMapClient{},
		},
	}

	dc.HTTPTool.JSONUnmarshal(c, options)
	if "" == options.Namespace || "" == options.DeploymentName {
		dc.HTTPTool.WriteErr(c, "参数无效")
	}
	rs, err := dc.CRDDependence.GetResources(options.Namespace, options.DeploymentName)
	if err != nil {
		dc.HTTPTool.WriteErr(c, err)
	} else {
		dc.HTTPTool.WriteOk(c, rs)
	}
}
