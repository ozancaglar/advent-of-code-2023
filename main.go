package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/ozancaglar/advent-of-code-2023/day1"
)

type day struct {
	Name     string
	Function func()
}

func main() {
	days := []day{
		{Name: "1", Function: day1.Solve},
	}

	templates := &promptui.SelectTemplates{
		Active:   "\U000025CF {{ .Name | blue }}",
		Inactive: "\U000025CB {{ .Name | red }}",
		Selected: "{{ .Name | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Which day would you like the solution to?",
		Items:     days,
		Templates: templates,
		Size:      8,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	days[i].Function()
}
