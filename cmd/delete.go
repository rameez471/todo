/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"todo/schema"
	"strconv"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Permanently delete a task from the list.",
	Run: runDelete,
}

func runDelete(cmd *cobra.Command, args []string) {
	items, err := schema.ReadItems(datafile)
	if err !=  nil{
		log.Fatal(err)
	}
	index, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	for i, _ := range items {
		if items[i].Position == index {
			items = schema.DeleteItem(items, i)
			break
		}
	}
	err = schema.WriteItems(datafile, items)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tak(s) deleted!")
}


func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
