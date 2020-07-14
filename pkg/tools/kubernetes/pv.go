package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PvClient struct
type PvClient struct{}

// Get pv by pvName
func (p *PvClient) Get(pvName string) (*v1.PersistentVolume, error) {
	pv, err := NewKubeClient().CoreV1().PersistentVolumes().Get(pvName, meta_v1.GetOptions{})
	if err != nil {
		logger.Errorf("Get pvc by pvcName error:%s,pvName:%s", err, pvName)
		return nil, err
	}
	return pv, nil
}
