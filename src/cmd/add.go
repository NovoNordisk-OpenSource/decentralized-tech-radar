/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <name> <ring> <quadrant> <isNew> <moved> <description>",
	Short: "adds a new line to a csv file",
	Long: `This command adds a new line to a csv file. The csv file must be in the following format:
	name,ring,quadrant,isNew,moved,description
	`,
	Args: cobra.MinimumNArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
