/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"fmt"
	"todo/schema"
	"github.com/spf13/cobra"
)

var priority int
// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command add a task into the list",
	Run: runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	items := [] schema.Item{}
	items, err := schema.ReadItems(datafile)
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range args {
		item := schema.Item{X: x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	err = schema.WriteItems(datafile, items)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Task added to the list!!")
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// IntVarP(p *int, name, shorthand string, value int, usage string)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Add priority to a Task\n 1: High\n 2: Medium\n 3: Low\n")
}
