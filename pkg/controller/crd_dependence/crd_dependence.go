package crd_dependence


import (
	"github.com/gin-gonic/gin"
	"resource/pkg/service/crd_dependence"
	"resource/pkg/tools/kubernetes"
)

type CRDDependenceOptions struct {
	DeploymentName string `json:"deploymentName"`
	Namespace      string `json:"namespace"`
}

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

	dc.HttpTool.JsonUnmarshal(c, options)
	if "" == options.Namespace || "" == options.DeploymentName {
		dc.HttpTool.WriteErr(c, "参数无效")
	}
	rs ,err := dc.CRDDependence.GetResources(options.Namespace, options.DeploymentName)
	if err != nil{
		dc.HttpTool.WriteErr(c,err)
	}else {
		dc.HttpTool.WriteOk(c,rs)
	}
}