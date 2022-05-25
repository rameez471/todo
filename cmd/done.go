/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"fmt"
	"todo/schema"
	"strconv"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark the completed tasks as done.",
	Run: runDone,
}

func runDone(cmd *cobra.Command, args []string) {
	items, err := schema.ReadItems(datafile)
	if err != nil{
		log.Fatal(err)
	}
	for _,x  := range args {
		index, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal("Please enter position number of Task to be marked as Done")
		}
		items[index-1].Done = true
		err = schema.WriteItems(datafile, items)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println("Tasks marked as Done!!")
	}
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
