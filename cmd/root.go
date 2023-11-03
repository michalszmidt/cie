/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"
	"fmt"
	"os"

	"github.com/michalszmidt/cie/customIO"
	"github.com/michalszmidt/cie/processing"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cie",
	Short: "Modify ICal calendar from cli",
	Long:  `Modify ICal calendar from cli`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: processingInit,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cie.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("path", "p",  "./cal.ics", "Path to the ics file")
	rootCmd.Flags().StringP("out", "o",  "./new.ics" ,"Path to out file")
	rootCmd.Flags().StringP("config", "c",  "./config.yml" ,"Path to config yml file")				
}

func processingInit(cmd *cobra.Command, args []string) {
	path, _ := cmd.Flags().GetString("path")
	out, _ := cmd.Flags().GetString("out")
	config, _ := cmd.Flags().GetString("config")
	cie := customIO.PathToYaml(config)
	counter := processing.Process(&path,&out,&cie)
	fmt.Println("Removed events: ", counter)
}
