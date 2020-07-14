package types

import (
	"resource/pkg/service"
)

// AppDependence 无状态应用获取依赖资源接口
type AppDependence interface {
	GetResources(deploymentName, namespace string) (*service.ResAppDependence, error)
}

// CRDDependence 有状态应用获取依赖资源接口
type CRDDependence interface {
	GetResources(namespace, stsName string) (*service.ResAppDependence, error)
}
