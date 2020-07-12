package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodClinet struct{}

// ListByLabels 根据标签获取pod列表
func (pc *PodClinet) ListByLabels(namespace string, listOptions meta_v1.ListOptions) (*v1.PodList, error) {
	podList, err := NewKubeClient().CoreV1().Pods(namespace).List(listOptions)
	if err != nil {
		logger.Errorf("Get podLost by  labels error:%s , namespace=%s", err, namespace)
		return nil, err
	}
	return podList, err
}
