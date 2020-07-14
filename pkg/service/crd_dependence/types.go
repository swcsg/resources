package crd_dependence

import (
	"resource/pkg/infoer"

	"github.com/maxwell92/log"
)

var logger = log.Log

// CRDDependence 封装 infoer接口
type CRDDependence struct {
	StatefulSetClient  infoer.StatefulSetClient
	PodClient          infoer.PodClient
	ServiceClient      infoer.ServiceClient
	PvcClient          infoer.PvcClient
	PvClient           infoer.PvClient
	ConfigMapClient    infoer.ConfigMapClient
	SecretClient       infoer.SecretClient
	StorageClassClient infoer.StorageClassClient
}
