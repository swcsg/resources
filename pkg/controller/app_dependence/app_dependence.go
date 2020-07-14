package app_dependence

import (
	"resource/pkg/service/app_dependence"
	"resource/pkg/tools/kubernetes"

	"github.com/gin-gonic/gin"
)

// AppDependenceOptions 有状态应用依赖资源请求结构体
type AppDependenceOptions struct {
	DeploymentName string `json:"deploymentName"`
	Namespace      string `json:"namespace"`
}

// AppDependence 无状态应用应用依赖资源控制层
func AppDependence(c *gin.Context) {

	var options = &AppDependenceOptions{}

	var dc = &AppDependenceController{
		AppDependence: &app_dependence.AppDependence{
			DeploymentClient:   &kubernetes.DeploymnetClient{},
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
		dc.HTTPTool.WriteErr(c, "Invalid argument")
	}
	rs, err := dc.AppDependence.GetResources(options.DeploymentName, options.Namespace)
	if err != nil {
		dc.HTTPTool.WriteErr(c, err)
	} else {
		dc.HTTPTool.WriteOk(c, rs)
	}

}
