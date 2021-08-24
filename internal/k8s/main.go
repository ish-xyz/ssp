package k8s

import (
	"log"

	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

func NewClient(configPath string) (*kubernetes.Clientset, error) {

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Panicln("failed to create K8s config")
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panicln("Failed to create K8s clientset")
		return nil, err
	}

	return clientset, nil
}
