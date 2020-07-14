package infoer

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	sc_v1beta1 "k8s.io/api/storage/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeploymentClient deployment 客户端接口
type DeploymentClient interface {
	Get(namespace, deploymentName string) (*v1beta1.Deployment, error)
}

// PodClient pod 客户端接口
type PodClient interface {
	ListByLabels(namespace string, listOptions meta_v1.ListOptions) (*v1.PodList, error)
}

// ServiceClient svc 客户端接口
type ServiceClient interface {
	List(namespace string, listOptions meta_v1.ListOptions) (*v1.ServiceList, error)
}

// PvcClient pvc 客户端接口
type PvcClient interface {
	Get(namespace, pvcName string) (*v1.PersistentVolumeClaim, error)
}

// PvClient  pv 客户端接口
type PvClient interface {
	Get(pvName string) (*v1.PersistentVolume, error)
}

// ConfigMapClient configmap 客户端接口
type ConfigMapClient interface {
	Get(namespace, name string) (*v1.ConfigMap, error)
}

// SecretClient secret 客户端接口
type SecretClient interface {
	Get(namespace, name string) (*v1.Secret, error)
}

// StorageClassClient sc 客户端接口
type StorageClassClient interface {
	Get(name string) (*sc_v1beta1.StorageClass, error)
}

// StatefulSetClient sts 客户端接口
type StatefulSetClient interface {
	Get(namespace, name string) (*appsv1.StatefulSet, error)
}
