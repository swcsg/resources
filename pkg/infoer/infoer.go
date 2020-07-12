package infoer

import (
	v1 "k8s.io/api/core/v1"
	appsv1"k8s.io/api/apps/v1"
	"k8s.io/api/extensions/v1beta1"
	sc_v1beta1 "k8s.io/api/storage/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentClient interface {
	Get(namespace, deploymentName string) (*v1beta1.Deployment, error)
}

type PodClient interface {
	ListByLabels(namespace string, listOptions meta_v1.ListOptions) (*v1.PodList, error)
}

type ServiceClient interface {
	List(namespace string, listOptions meta_v1.ListOptions) (*v1.ServiceList, error)
}

type PvcClient interface {
	Get(namespace, pvcName string) (*v1.PersistentVolumeClaim, error)
}

type PvClient interface {
	Get(pvName string) (*v1.PersistentVolume, error)
}

type ConfigMapClient interface {
	Get(namespace, name string) (*v1.ConfigMap, error)
}

type SecretClient interface {
	Get(namespace, name string) (*v1.Secret, error)
}

type StorageClassClient interface {
	Get(name string) (*sc_v1beta1.StorageClass, error)
}

type StatefulSetClient interface {
	Get(namespace,name string)(*appsv1.StatefulSet,error)
}
