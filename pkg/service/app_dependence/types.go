package app_dependence

import (
	"resource/pkg/infoer"

	"github.com/maxwell92/log"
)

var logger = log.Log

// AppDependence 封装 infoer接口
type AppDependence struct {
	DeploymentClient   infoer.DeploymentClient
	PodClient          infoer.PodClient
	ServiceClient      infoer.ServiceClient
	PvcClient          infoer.PvcClient
	PvClient           infoer.PvClient
	ConfigMapClient    infoer.ConfigMapClient
	SecretClient       infoer.SecretClient
	StorageClassClient infoer.StorageClassClient
}
