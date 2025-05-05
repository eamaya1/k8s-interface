package main

import (
	"fmt"
)

func (k kubeInterface) showWelcomeView() string {
	s := "Welcome to your personal cluster manager! \nPlease choose from the following choices:"
	for idx, val := range k.choices {
		s += fmt.Sprintf("\n%d : %s", idx, val)
	}

	s += "\nOr press ctrl + C to exit"
	return s
}
