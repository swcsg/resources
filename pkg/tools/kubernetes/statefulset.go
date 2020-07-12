package kubernetes

import (
	"k8s.io/api/apps/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatefulSet struct {}

func (s *StatefulSet)Get(namespace ,name string)(*v1.StatefulSet,error){
	sts,err := NewKubeClient().AppsV1().StatefulSets(namespace).Get(name,meta_v1.GetOptions{})
	if err != nil{
		logger.Errorf("Get StatefulSet error:%s,namespace:%s,name:%s",err,namespace,name)
		return nil,err
	}
	return sts,err
}
