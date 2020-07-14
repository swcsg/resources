package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PvcClient struct
type PvcClient struct{}

// Get get pvc
func (p *PvcClient) Get(namespace, name string) (*v1.PersistentVolumeClaim, error) {
	pvc, err := NewKubeClient().CoreV1().PersistentVolumeClaims(namespace).Get(name, meta_v1.GetOptions{})
	if err != nil {
		logger.Errorf("Get pvc by pvcName error:%s,pvcName:%s,namespace:%s", err, name, namespace)
		return nil, err
	}
	return pvc, nil
}
