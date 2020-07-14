package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceClient struct
type ServiceClient struct{}

// List get serviceList by listOptions
func (s *ServiceClient) List(namespace string, listOptions meta_v1.ListOptions) (*v1.ServiceList, error) {
	svcList, err := NewKubeClient().CoreV1().Services(namespace).List(listOptions)
	if err != nil {
		logger.Errorf("Get service by listoptions error:%s , namespace:%s", err, namespace)
		return nil, err
	}
	return svcList, err
}
