package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigMapClient struct{}

func (c *ConfigMapClient) Get(namespace, name string) (*v1.ConfigMap, error) {
	svcList, err := NewKubeClient().CoreV1().ConfigMaps(namespace).Get(name, meta_v1.GetOptions{})
	if err != nil {
		logger.Errorf("Get configMap by name error:%s , namespace:%s", err, namespace)
		return nil, err
	}
	return svcList, err
}
