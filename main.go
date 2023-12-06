package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/ozancaglar/advent-of-code-2023/day1"
	"github.com/ozancaglar/advent-of-code-2023/day2"
	"github.com/ozancaglar/advent-of-code-2023/day3"
	"github.com/ozancaglar/advent-of-code-2023/day4"
	"github.com/ozancaglar/advent-of-code-2023/day5"
	"github.com/ozancaglar/advent-of-code-2023/day6"
)

type day struct {
	Day      string
	Function func(filename string)
}

func main() {

	days := []day{
		{Day: "1", Function: day1.Solve},
		{Day: "2", Function: day2.Solve},
		{Day: "3", Function: day3.Solve},
		{Day: "4", Function: day4.Solve},
		{Day: "5", Function: day5.Solve},
		{Day: "6", Function: day6.Solve},
	}

	templates := &promptui.SelectTemplates{
		Active:   "\U000025CF {{ .Day | blue }}",
		Inactive: "\U000025CB {{ .Day | red }}",
		Selected: "{{ .Day | red | cyan }}",
	}

	searcher := func(input string, index int) bool {
		day := days[index]
		name := strings.Replace(strings.ToLower(day.Day), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Which day would you like the solution to?",
		Items:     days,
		Templates: templates,
		Size:      8,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	days[i].Function(fmt.Sprintf("day%s/input.txt", strconv.Itoa(i+1)))
}
