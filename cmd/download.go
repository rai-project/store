package cmd

import (
	"errors"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

var (
	outputFile string
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use: "download",
	Aliases: []string{
		"get",
	},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if outputFile == "" {
			return errors.New("Invalid output file")
		}
		pp.Println(args)
		return nil
	},
}

func init() {
	downloadCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "output file")
}
