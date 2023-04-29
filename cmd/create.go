/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/lunarisnia/todo-cli/data"
	"github.com/lunarisnia/todo-cli/prompt"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new todo list",
	Long:  `Create a new todo list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		todoTitlePc := prompt.PromptContent{
			ErrorMessage: "title is invalid",
			Label:        "Title",
		}
		todoDescPc := prompt.PromptContent{
			ErrorMessage: "description is invalid",
			Label:        "Description",
		}

		title := prompt.PromptGetInput(todoTitlePc, false)
		description := prompt.PromptGetInput(todoDescPc, true)
		data.InsertTodo(title, description, false)

		fmt.Fprintf(cmd.OutOrStdout(), "A Todo has been added with title: %v \n", title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
