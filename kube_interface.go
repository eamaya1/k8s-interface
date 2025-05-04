package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Defining our model
type kubeInterface struct {
	choices  []string
	cursor   int
	selected int // Choosing 1 at a time, the selected would be the index of the choices list
}

func newkubeInterface(choices []string, cursor int, selected int) kubeInterface {
	return kubeInterface{
		choices:  choices,
		cursor:   cursor,
		selected: selected,
	}
}

func (m kubeInterface) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now"
	return nil
}

func (k kubeInterface) View() string {
	s := "****Welcome to your cluster manager"
	return s
}

func (k kubeInterface) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return k, tea.Quit
		}
	}
	return k, nil
}
