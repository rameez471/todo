/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"todo/schema"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command add a task into the list",
	Run: runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	items := [] schema.Item{}
	items, err := schema.ReadItems("/home/rameez/.tododata.json")
	if err != nil {
		fmt.Println(err)
	}
	for _, x := range args {
		item := schema.Item{X: x}
		items = append(items, item)
	}
	fmt.Println(items)
	err = schema.WriteItems("/home/rameez/.tododata.json", items)
	if err != nil{
		fmt.Println(err)
	}

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
}