package kubernetes

import (
	"k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymnetClient struct{}

func (d *DeploymnetClient) Get(namespace, deploymentName string) (*v1beta1.Deployment, error) {
	dp, err := NewKubeClient().ExtensionsV1beta1().Deployments(namespace).Get(deploymentName, v1.GetOptions{})
	if err != nil {
		logger.Errorf("Get deployment error:%s , namespace:%s , deploymentName:%s", err, namespace, deploymentName)
	}
	return dp, nil
}
