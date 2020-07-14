package service

import (
	apps_v1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	sc_v1beta1 "k8s.io/api/storage/v1beta1"
)

// ResAppDependence 请求返回结构体
type ResAppDependence struct {
	Deployment   *v1beta1.Deployment
	StatefulSet  *apps_v1.StatefulSet
	Service      *v1.ServiceList
	Pod          *v1.PodList
	ConfigMap    []*v1.ConfigMap
	Secret       []*v1.Secret
	Pvc          []*v1.PersistentVolumeClaim
	Pv           []*v1.PersistentVolume
	StorageClass []*sc_v1beta1.StorageClass
}
