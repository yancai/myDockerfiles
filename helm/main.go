package main

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/helm/pkg/helm"
	helm_env "k8s.io/helm/pkg/helm/environment"
	"k8s.io/helm/pkg/helm/portforwarder"
	"k8s.io/helm/pkg/kube"
	"log"
)

var (
	settings helm_env.EnvSettings
)

// configForContext creates a Kubernetes REST client configuration for a given kubeconfig context.
func configForContext(context string) (*rest.Config, error) {
	config, err := kube.GetConfig(context).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("could not get Kubernetes config for context %q: %s", context, err)
	}
	return config, nil
}

// getKubeClient creates a Kubernetes config and client for a given kubeconfig context.
func getKubeClient(context string) (*rest.Config, kubernetes.Interface, error) {
	config, err := configForContext(context)
	if err != nil {
		return nil, nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get Kubernetes client: %s", err)
	}
	return config, client, nil
}

func debug(format string, args ...interface{}) {
	if settings.Debug {
		format = fmt.Sprintf("[debug] %s\n", format)
		fmt.Printf(format, args...)
	}
}

func main() {
	fmt.Println("hello")
	if settings.TillerHost == "" {
		config, client, err := getKubeClient(settings.KubeContext)
		if err != nil {
			log.Fatal(err)
		}

		tunnel, err := portforwarder.New(settings.TillerNamespace, client, config)
		if err != nil {
			log.Fatal(err)
		}

		settings.TillerHost = fmt.Sprintf("127.0.0.1:%d", tunnel.Local)
		debug("Created tunnel using local port: '%d'\n", tunnel.Local)
	}
	fmt.Println(settings.TillerHost)

	cli := helm.NewClient()
	result, err := cli.GetVersion()
	if err != nil {
		log.Fatal("get version error: ", err)
	}

	fmt.Println(result)
}
