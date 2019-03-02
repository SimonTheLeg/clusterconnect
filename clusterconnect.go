package clusterconnect

import (
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	flag "github.com/spf13/pflag"
)

// ConnectToCluster connects to a Kubernetes cluster. It can connect either via a local kubeconfig file or from within a Pod itself
// It switches between the two by registering a --local flag
func ConnectToCluster() (*kubernetes.Clientset, error) {
	var err error
	var config *rest.Config
	var local bool

	flag.BoolVar(&local, "local", false, "Specify whether to use the local kubeconfig or an in-cluster one")

	flag.Parse()

	if local == true {
		// Create the kubeconfig from the local kubeconfig file
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		log.Println("Using kubeconfig: ", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		// Generate a kubeconfig from within the cluster
		log.Println("Using in-cluster config")
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
