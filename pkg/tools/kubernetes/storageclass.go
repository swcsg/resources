package kubernetes

import (
	sc_v1beta1 "k8s.io/api/storage/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StroageClassClient struct
type StroageClassClient struct{}

// Get StroageClassClient
func (s *StroageClassClient) Get(name string) (*sc_v1beta1.StorageClass, error) {
	sc, err := NewKubeClient().StorageV1beta1().StorageClasses().Get(name, meta_v1.GetOptions{})
	if err != nil {
		logger.Errorf("Get stroageClass by name error:%s , name:%s", err, name)
		return nil, err
	}
	return sc, err
}
