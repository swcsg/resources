package kubernetes

import (
	"flag"
	"sync"

	"github.com/maxwell92/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var logger = log.Log

var once sync.Once

var kubeClient *kubernetes.Clientset

// NewKubeClient new kubernetes clientset
func NewKubeClient() *kubernetes.Clientset {
	once.Do(func() {
		kubeClient = kubeClientInit()
	})
	return kubeClient
}

// kubeClientInit clientset init , return a kubernetes clientset
func kubeClientInit() *kubernetes.Clientset {

	kubeconfig := flag.String("kubeconfig", "/root/.kube/config", "absolute path to the kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// 创建client set
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
