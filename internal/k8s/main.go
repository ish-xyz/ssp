package k8s

import (
	"log"

	clientcmd "k8s.io/client-go/1.5/tools/clientcmd"
	kubernetes "k8s.io/client-go/kubernetes"
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
