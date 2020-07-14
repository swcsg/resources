package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodClinet struct
type PodClinet struct{}

// ListByLabels get podList by listOptions
func (pc *PodClinet) ListByLabels(namespace string, listOptions meta_v1.ListOptions) (*v1.PodList, error) {
	podList, err := NewKubeClient().CoreV1().Pods(namespace).List(listOptions)
	if err != nil {
		logger.Errorf("Get podLost by  labels error:%s , namespace=%s", err, namespace)
		return nil, err
	}
	return podList, err
}
