package app_dependence

import (
	"github.com/gin-gonic/gin"
	"resource/pkg/service/app_dependence"
	"resource/pkg/tools/kubernetes"
)

type AppDependenceOptions struct {
	DeploymentName string `json:"deploymentName"`
	Namespace      string `json:"namespace"`
}

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

	dc.HttpTool.JsonUnmarshal(c, options)
	if "" == options.Namespace || "" == options.DeploymentName {
		dc.HttpTool.WriteErr(c, "参数无效")
	}
	rs ,err := dc.AppDependence.GetResources(options.DeploymentName, options.Namespace)
	if err != nil{
		dc.HttpTool.WriteErr(c,err)
	}else {
		dc.HttpTool.WriteOk(c,rs)
	}


}
