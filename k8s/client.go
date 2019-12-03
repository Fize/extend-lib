package k8s

import (
	"context"
	"encoding/json"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	clientcmdlatest "k8s.io/client-go/tools/clientcmd/api/latest"
	clientcmdapiv1 "k8s.io/client-go/tools/clientcmd/api/v1"
)

// GetK8sClient return a k8s client, v is usually a SchemeGroupVersion of a API
func GetK8sClient(ctx context.Context, restClient bool, master, kubeconfig, apiPath string, v interface{}) (
	*kubernetes.Clientset, *rest.RESTClient, error) {
	if kubeConfig == "" {
		restConfig, err = inCluster()
		if err != nil {
			return nil, nil, err
		}
	} else {
		restConfig, err = outCluster(master, kubeconfig)
		if err != nil {
			return nil, nil, err
		}
	}
	if restClient {
		restConfig.GroupVersion = v
		restConfig.APIPath = apiPath
		restConfig.ContentType = runtime.ContentTypeJSON
		restConfig.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
		restclient, err := rest.RESTClientFor(restConfig)
		return nil, restclient, err
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	return clientset, nil, err
}

// if use inCluster，configure service account for deployment
func inCluster() (*rest.Config, error) {
	return rest.InClusterConfig()
}

// if use outCluster, specified kubeconfig file
func outCluster(master, config string) (*rest.Config, error) {
	configByte, err := yaml.ToJSON([]byte(config))
	if err != nil {
		return nil, err
	}
	configV1 := clientcmdapiv1.Config{}
	err = json.Unmarshal(configByte, &configV1)
	if err != nil {
		return nil, err
	}
	configObject, err := clientcmdlatest.Scheme.ConvertToVersion(&configV1, clientcmdapi.SchemeGroupVersion)
	if err != nil {
		return nil, err
	}
	configInternal := configObject.(*clientcmdapi.Config)

	return clientcmd.NewDefaultClientConfig(*configInternal, &clientcmd.ConfigOverrides{
		ClusterDefaults: clientcmdapi.Cluster{Server: master},
	}).ClientConfig()
}
