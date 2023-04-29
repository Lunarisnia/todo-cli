/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/lunarisnia/todo-cli/data"
	"github.com/lunarisnia/todo-cli/prompt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var showEverything bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all todo list",
	Long:  `Show all todo list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		todos := data.ReadAllTodos(showEverything)

		// TODO: add pagination
		// TODO: Modify the cursor icon
		// TODO: Change the ongoing text color to red and done to green
		templates := promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "\U0001F336 {{ .Title | cyan }}",
			Inactive: "  {{ .Title | cyan }}",
			Selected: "\U0001F336 {{ .Title | red | cyan }}",
			Details: `
----------- Details -----------
{{ "Title:" | faint }}  {{ .Title }}
{{ "Description:" | faint }}  {{ .Description }}
{{ "Status:" | faint }}  {{ if .Status }} {{ "✅" }} {{ else }} {{ "❌" }} {{ end }}
`,
		}

		sc := prompt.SelectContent{
			Label:     "Your Todo List",
			Items:     todos,
			Templates: &templates,
		}
		prompt.PromptSelectContent(&sc)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().BoolVarP(&showEverything, "everything", "e", false, "Show everything including finished TODO.")
}
