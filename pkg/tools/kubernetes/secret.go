package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SecretClient struct{}

func (s *SecretClient) Get(namespace, name string) (*v1.Secret, error) {
	secret, err := NewKubeClient().CoreV1().Secrets(namespace).Get(name, meta_v1.GetOptions{})
	if err != nil {
		logger.Errorf("Get secret by name error:%s,pvName:%s", err, name)
		return nil, err
	}
	return secret, nil
}
