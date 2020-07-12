package crd_dependence

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"resource/pkg/service"
)

func (crd *CRDDependence) GetResources(namespace, name string) (*service.ResAppDependence, error) {

	var res = &service.ResAppDependence{}
	// sts
	sts, err := crd.StatefulSetClient.Get(namespace, name)
	if err != nil {
		return nil, err
	}
	res.StatefulSet = sts
	selector, err := meta_v1.LabelSelectorAsSelector(sts.Spec.Selector)
	if err != nil {
		logger.Errorf("Get lables  by deployment err :%s", err)
		return nil, err

	}
	listOptions := meta_v1.ListOptions{LabelSelector: selector.String()}
	logger.Infoln("listOptions:%s", listOptions)

	// podList
	podList, err := crd.PodClient.ListByLabels(namespace, listOptions)
	if err != nil {
		// pod不存在，对应依赖的资源无意义
		return nil, err
	}
	res.Pod = podList

	// service
	svcList, err := crd.ServiceClient.List(namespace, listOptions)
	if err == nil {
		res.Service = svcList
	}

	volumes := sts.Spec.Template.Spec.Volumes


	for _, volume := range volumes { //可能多个资源挂载
		// secret
		if volume.Secret != nil {
			secret, err := crd.SecretClient.Get(namespace, volume.Secret.SecretName)
			if err == nil {
				res.Secret = append(res.Secret, secret)
			}
		}
		// configMap
		if volume.ConfigMap != nil {
			configMap, err := crd.ConfigMapClient.Get(namespace, volume.ConfigMap.Name)
			if err == nil {
				res.ConfigMap = append(res.ConfigMap, configMap)
			}
		}
	}
	// pvc
	persistentVolumeClaim := sts.Spec.VolumeClaimTemplates
	if persistentVolumeClaim != nil {
		for _, p := range persistentVolumeClaim {
			pvc, err := crd.PvcClient.Get(namespace, p.Name)
			if err != nil {
				// pvc 不存在，pv，sc 无意义
				continue
			}
			res.Pvc = append(res.Pvc, pvc)
			// sc
			if pvc.Spec.StorageClassName == nil {
				sc, err := crd.StorageClassClient.Get(*pvc.Spec.StorageClassName)
				if err == nil {
					res.StorageClass = append(res.StorageClass, sc)
				}
			}
			// pv
			pv, err := crd.PvClient.Get(pvc.Spec.VolumeName)
			if err == nil {
				res.Pv = append(res.Pv, pv)
			}
		}
	}
	return res, err
}
