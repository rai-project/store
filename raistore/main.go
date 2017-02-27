package main

import (
	"fmt"
	"os"

	"github.com/rai-project/config"
	"github.com/rai-project/store/cmd"
	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var root = &cobra.Command{
	Use: "raistore",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the root.
func main() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	root.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	root.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmd.Init(root)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		config.ConfigFileName = cfgFile
	}
	config.Init()
}
