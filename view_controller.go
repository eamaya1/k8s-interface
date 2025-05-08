package main

import tea "github.com/charmbracelet/bubbletea"

type viewController struct {
	activeView tea.Model
}

func newViewController() viewController {
	choices := []string{"pods", "services", "deployments"}
	return viewController{
		activeView: newKubeInterface(choices, 0, -1),
	}
}

func (vc viewController) Init() tea.Cmd {
	return nil
}
func (vc viewController) View() string {
	return vc.activeView.View()
}

func (vc viewController) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//Create view logic,
	return vc.activeView.Update(msg)
}
