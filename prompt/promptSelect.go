package prompt

import (
	"fmt"

	"github.com/lunarisnia/todo-cli/data"
	"github.com/manifoldco/promptui"
)

type SelectContent struct {
	Label     string
	Items     []data.Todo
	Templates *promptui.SelectTemplates
}

func PromptSelectContent(sc *SelectContent) {
	prompt := promptui.Select{
		Label:     sc.Label,
		Items:     sc.Items,
		Templates: sc.Templates,
	}

	_, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
}
