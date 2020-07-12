package app_dependence

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"resource/pkg/service"
)

func (d *AppDependence) GetResources(deploymentName, namespace string) (*service.ResAppDependence,error){

	var res = &service.ResAppDependence{}
	// deployment
	dp, err := d.DeploymentClient.Get(namespace, deploymentName)
	if err != nil {
		return nil,err
	}
	res.Deployment = dp

	selector, err := meta_v1.LabelSelectorAsSelector(dp.Spec.Selector)
	if err != nil {
		logger.Errorf("Get lables  by deployment err :%s", err)
		return nil,err

	}
	listOptions :=  meta_v1.ListOptions{LabelSelector: selector.String()}
	logger.Infoln("listOptions:%s",listOptions)

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

	volumes := dp.Spec.Template.Spec.Volumes

	for _, volume := range volumes {
		// secret
		if volume.Secret != nil {
			secret, err := d.SecretClient.Get(namespace, volume.Secret.SecretName)
			if err == nil {
				res.Secret = append(res.Secret,secret)
			}
		}
		// configMap
		if volume.ConfigMap != nil {
			configMap, err := d.ConfigMapClient.Get(namespace, volume.ConfigMap.Name)
			if err == nil {
				res.ConfigMap = append(res.ConfigMap,configMap)
			}
		}
		// pvc
		if volume.PersistentVolumeClaim != nil {
			pvc, err := d.PvcClient.Get(namespace, volume.PersistentVolumeClaim.ClaimName)
			if err != nil {
				// pvc 不存在，pv，sc 无意义
				continue
			}
			res.Pvc = append(res.Pvc,pvc)
			// sc
			if pvc.Spec.StorageClassName == nil {
				sc, err := d.StorageClassClient.Get(*pvc.Spec.StorageClassName)
				if err == nil {
					res.StorageClass = append(res.StorageClass,sc)
				}
			}
			// pv
			pv, err := d.PvClient.Get(pvc.Spec.VolumeName)
			if err == nil{
				res.Pv = append(res.Pv,pv)
			}

		}
	}
	return res,nil

}