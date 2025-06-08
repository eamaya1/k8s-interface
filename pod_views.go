package main

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getPods(clientset *kubernetes.Clientset) *corev1.PodList {
	pods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Printf("There was an error getting the pods: %s", err)
		os.Exit(1)
	}

	// Type v1.PodList, access it with pods.Items[].Name
	return pods
}

// Defining our view model for any new screens we need to show
type podView struct {
	view      string
	clientset *kubernetes.Clientset
}

func newPodView(clientset *kubernetes.Clientset) podView {
	s := "You are now looking at the Pods screen"
	return podView{
		view:      s,
		clientset: clientset,
	}
}

func (v podView) Init() tea.Cmd {
	return nil
}

// There is a problem with clicking 0: pods, pointers may have been lost. Investigate that with the debugger if possible
func (v podView) View() string {
	pods := getPods(v.clientset)
	numPods := len(pods.Items)
	s := fmt.Sprintf("Here are the number of pods in your cluster: %d\n Here are the pods in your cluster:", numPods)

	for _, pod := range pods.Items {
		s += pod.Name + "\n"
	}

	return s
}

func (v podView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return v, tea.Quit
		case "backspace":
			return v, func() tea.Msg {
				choices := []string{"pods", "services", "deployments"}
				return switchViewMessage{newView: &kubeInterface{choices, 0, -1, v.clientset}}
			} //Add the logic for this
		}
	}

	return v, nil
}
