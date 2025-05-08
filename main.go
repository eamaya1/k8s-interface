package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	//choices := []string{"pods", "services", "deployments"}
	p := tea.NewProgram(
		//newKubeInterface(choices, 0, -1),
		newViewController(),
	)

	if err := p.Start(); err != nil {
		panic(err)
	}
}
