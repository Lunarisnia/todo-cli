package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type SelectContent struct {
	Label     string
	Items     []interface{}
	Templates *promptui.SelectTemplates
}

func PromptSelectContent(sc *SelectContent) (int, string) {
	prompt := promptui.Select{
		Label:        sc.Label,
		Items:        sc.Items[0],
		Templates:    sc.Templates,
		HideSelected: true,
	}

	index, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 0, ""
	}

	return index, result
}
