/*
Copyright © 2022 rameez471

*/
package cmd

import (
	"os"
	"log"
	"github.com/spf13/cobra"
	"github.com/mitchellh/go-homedir"
)

var datafile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple cli application made in Go, with cobra package",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	homedir, err := homedir.Dir()
	if err != nil{
		log.Fatal(err)
	}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//  StringVar(p *string, name string, value string, usage string)
	rootCmd.PersistentFlags().StringVar(&datafile, "datafile", homedir+string(os.PathSeparator)+".tododata.json", "location of storage")
}


