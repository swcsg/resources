package app_dependence

import (
	"github.com/maxwell92/log"
	"resource/pkg/infoer"
)

var logger = log.Log

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


