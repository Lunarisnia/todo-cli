package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type SelectContent struct {
	Label string
	Items []string
}

func PromptSelectContent(sc *SelectContent) {
	prompt := promptui.Select{
		Label: sc.Label,
		Items: sc.Items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	fmt.Printf("You choose %q\n", result)
}
