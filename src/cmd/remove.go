package cmd

import "github.com/spf13/cobra"

var remCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(remCmd)
}
