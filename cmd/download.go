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
