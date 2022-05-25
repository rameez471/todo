/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"todo/schema"
	"log"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the tasks available in the list",
	Run: runList,
}

func runList(cmd *cobra.Command, args []string) {
	items := []schema.Item{}
	items, err := schema.ReadItems("/home/rameez/.tododata.json")
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range items {
		fmt.Println(x.Pretty())
	}
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
}
