package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Defining our view model for any new screens we need to show
type podView struct {
	view string
}

func newPodView(view string) podView {
	s := "You are now looking at the Pods screen"
	return podView{
		view: s,
	}
}

func (v podView) Init() tea.Cmd {
	return nil
}

func (v podView) View() string {
	v.view = "Here are the number of pods in your cluster:"
	v.view += "\nPress ctrl+c to exit or backspace to return to main menu"
	return v.view
}

func (v podView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return v, tea.Quit
		case "backspace":
			return v, tea.Println("Returning to main screen") //Add the logic for this
		}
	}

	return v, nil
}
