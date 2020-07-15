package kubernetes

import (
	"flag"
	"sync"

	"github.com/maxwell92/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var logger = log.Log

var once sync.Once

var kubeClient *kubernetes.Clientset

// NewKubeClient new kubernetes clientset
func NewKubeClient() *kubernetes.Clientset {
	once.Do(func() {
		kubeClient = NewKubeClientInit()
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

// NewKubeClientInit TLS 认证方式
func NewKubeClientInit() *kubernetes.Clientset {
	// 若集群使用 TLS 认证方式，则默认读取集群内部 tokenFile 和 CAFile
	// tokenFile  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	// rootCAFile = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
	config, err := rest.InClusterConfig()
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
