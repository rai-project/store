package cmd

import "github.com/spf13/cobra"

// Init adds the toText and png commands to rootCmd
func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(uploadCmd)
}
