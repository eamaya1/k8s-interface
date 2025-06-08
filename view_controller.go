package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"k8s.io/client-go/kubernetes"
)

type viewController struct {
	kI         kubeInterface
	pv         podView
	activeView tea.Model
	clientset  *kubernetes.Clientset
}

type switchViewMessage struct {
	newView tea.Model
}

func newViewController(clientset *kubernetes.Clientset) viewController {
	choices := []string{"pods", "services", "deployments"}
	kI := newKubeInterface(choices, 0, -1, clientset)
	return viewController{
		kI:         kI,
		pv:         newPodView(clientset),
		activeView: &kI,
		clientset:  clientset,
	}
}

func (vc *viewController) Init() tea.Cmd {
	return nil
}
func (vc *viewController) View() string {
	return vc.activeView.View()
}

// Keeping it like this for now, with the caveat that we are currently creating a new view each time we go to a new page
// Of course, we'd want to not do this for efficiency, but to get a working prototype for now this will do.
func (vc *viewController) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//Create view logic,
	switch m := msg.(type) {
	case switchViewMessage:
		vc.activeView = m.newView
		return vc, nil

	default:
		var cmd tea.Cmd
		vc.activeView, cmd = vc.activeView.Update(msg)
		return vc, cmd
	}
}
