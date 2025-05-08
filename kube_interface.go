package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Defining our model
type kubeInterface struct {
	choices  []string
	cursor   int
	selected int // Choosing 1 at a time, the selected would be the index of the choices list
}

func newKubeInterface(choices []string, cursor int, selected int) kubeInterface {
	return kubeInterface{
		choices:  choices,
		cursor:   cursor,
		selected: selected,
	}
}

func (k kubeInterface) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now"
	return nil
}

// View, a function that renders the UI based on the data in the model.
// Used to render our UI. Of all the methods, the view is the simplest.
//
//		We look at the model in its current state and use it to return a string. That string is the UI
//	 Because the view describes the entire UI of your application, you don’t have to worry about redrawing logic and stuff like that. Bubble Tea takes care of it
func (k kubeInterface) View() string {
	return k.renderMainView()
}

func (k kubeInterface) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return k, tea.Quit
		case "up":
			if k.cursor > 0 {
				k.cursor--
				//return k, tea.Printf("Selected pods")
			}
		case "down":
			if k.cursor < len(k.choices)-1 {
				k.cursor++
				//fmt.Println(k.selected)
			}
		case "enter":
			// selected := k.cursor
			// fmt.Println(selected)
			return newPodView("").Update(msg) //tea.Cmd Would want this to execute whatever command the users wants at that specific cursor ex) at 1, show cluster services
		}
	}
	return k, nil
}

func (k kubeInterface) renderMainView() string {
	s := "Welcome to your personal cluster manager! \nPlease choose from the following choices:"
	for i, choice := range k.choices {
		cursor := " "
		if k.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("\n%s %d : %s", cursor, i, choice)
	}

	s += "\nOr press ctrl + C to exit"
	return s
}
