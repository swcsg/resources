package app_dependence

import (
	"resource/pkg/service"

	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetResources 获取有状态应用依赖资源对象的核心逻辑
func (d *AppDependence) GetResources(deploymentName, namespace string) (*service.ResAppDependence, error) {

	var res = &service.ResAppDependence{}
	// deployment
	dp, err := d.DeploymentClient.Get(namespace, deploymentName)
	if err != nil {
		return nil, err
	}
	res.Deployment = dp

	selector, err := meta_v1.LabelSelectorAsSelector(dp.Spec.Selector)
	if err != nil {
		logger.Errorf("Get lables  by deployment err :%s", err)
		return nil, err

	}
	// 获取标签
	listOptions := meta_v1.ListOptions{LabelSelector: selector.String()}

	// podList
	podList, err := d.PodClient.ListByLabels(namespace, listOptions)
	if err != nil {
		// pod不存在，对应依赖的资源无意义
		return nil, err
	}
	res.Pod = podList

	// service
	svcList, err := d.ServiceClient.List(namespace, listOptions)
	if err == nil {
		res.Service = svcList
	}
	if err != nil && errors.IsNotFound(err) {
		logger.Infoln("svc is not exist")
	}

	volumes := dp.Spec.Template.Spec.Volumes

	for _, volume := range volumes {
		// secret
		if volume.Secret != nil {
			secret, err := d.SecretClient.Get(namespace, volume.Secret.SecretName)
			if err == nil {
				res.Secret = append(res.Secret, secret)
			}
			if err != nil && errors.IsNotFound(err) {
				logger.Infoln("secret is not exist")
			}
		}
		// configMap
		if volume.ConfigMap != nil {
			configMap, err := d.ConfigMapClient.Get(namespace, volume.ConfigMap.Name)
			if err == nil {
				res.ConfigMap = append(res.ConfigMap, configMap)
			}
			if err != nil && errors.IsNotFound(err) {
				logger.Infoln("configmap is not exist")
			}
		}
		// pvc
		if volume.PersistentVolumeClaim != nil {
			pvc, err := d.PvcClient.Get(namespace, volume.PersistentVolumeClaim.ClaimName)
			if err != nil {
				// pvc 不存在，pv，sc 无意义
				logger.Infoln("pvc is not exist")
				continue
			}
			res.Pvc = append(res.Pvc, pvc)
			// sc
			if pvc.Spec.StorageClassName == nil {
				sc, err := d.StorageClassClient.Get(*pvc.Spec.StorageClassName)
				if err == nil {
					res.StorageClass = append(res.StorageClass, sc)
				}
				if err != nil && errors.IsNotFound(err) {
					logger.Infoln("sc is not exist")
				}
			}
			// pv
			pv, err := d.PvClient.Get(pvc.Spec.VolumeName)
			if err == nil {
				res.Pv = append(res.Pv, pv)
			}
			if err != nil && errors.IsNotFound(err) {
				logger.Infoln("pv is not exist")
			}

		}
	}
	return res, nil

}
