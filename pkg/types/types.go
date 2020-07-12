package types

import (
	"resource/pkg/service"
)

type AppDependence interface {
	GetResources(deploymentName, namespace string) (*service.ResAppDependence,error)
}

type CRDDependence interface {
	GetResources(namespace, stsName string)(*service.ResAppDependence,error)
}
