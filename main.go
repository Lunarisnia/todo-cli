/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/lunarisnia/todo-cli/cmd"
	"github.com/lunarisnia/todo-cli/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
