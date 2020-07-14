package kubernetes

import (
	v1 "k8s.io/api/rbac/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceAccountCient struct
type ServiceAccountCient struct{}

// GetClusterRoler get clusterRole
func (s *ServiceAccountCient) GetClusterRoler(name string) (*v1.ClusterRole, error) {
	sa, err := NewKubeClient().RbacV1().ClusterRoles().Get(name, meta_v1.GetOptions{})
	if err != nil {
		logger.Infoln("Get serviceAccount error:%s", err)
		return nil, err
	}
	return sa, nil
}

// GetRoler get role
func (s *ServiceAccountCient) GetRoler(name, namespace string) (*v1.Role, error) {
	sa, err := NewKubeClient().RbacV1().Roles(namespace).Get(name, meta_v1.GetOptions{})
	if err != nil {
		logger.Infoln("Get serviceAccount error:%s", err)
		return nil, err
	}
	return sa, nil
}
