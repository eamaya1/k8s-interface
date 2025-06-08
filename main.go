package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}

func setupKubeconfig(pathToConfig string) (*rest.Config, error) {

	kubeconfig, err := clientcmd.BuildConfigFromFlags("", pathToConfig)
	if err != nil {
		fmt.Printf("Something went wrong setting up the kubeconfig")
		os.Exit(1)
	}

	return kubeconfig, nil
}

func main() {
	//choices := []string{"pods", "services", "deployments"}
	kubeconfig, err := setupKubeconfig(goDotEnvVariable("KUBECONFIG"))
	if err != nil {
		fmt.Printf("Invalid path to config")
	}

	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		fmt.Printf("There was an error creating the clientset. Check kubeconfig")
		os.Exit(1)
	}

	// pods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), v1.ListOptions{})
	// if err != nil {
	// 	fmt.Printf("There was some error getting the pods from the namespace provided")
	// 	os.Exit(1)
	// }

	// fmt.Println("Here are the pods in the selected namespace")
	// for _, pod := range pods.Items {
	// 	fmt.Println(pod.Name)
	// }

	var v viewController = newViewController(clientset)
	vc := &v
	p := tea.NewProgram(
		//newKubeInterface(choices, 0, -1),
		vc,
	)

	if err := p.Start(); err != nil {
		panic(err)
	}
}
