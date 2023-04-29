package prompt

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type PromptContent struct {
	ErrorMessage string
	Label        string
}

func PromptGetInput(pc PromptContent, optional bool) string {
	validate := func(input string) error {
		if !optional && (input == "") {
			return errors.New(pc.ErrorMessage)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    pc.Label,
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
