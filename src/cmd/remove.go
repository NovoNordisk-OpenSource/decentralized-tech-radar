package cmd

import "github.com/spf13/cobra"

var remCmd = &cobra.Command{
	Use:   "remove <path/to/csvfile> <name> <quadrant>",
	Short: "Removes a line from a csv spec file",
	Long: `Removes the specified tool's line from the given csv specification file. The csv file's header must be in the following format:
		name,ring,quadrant,isNew,moved,description`,
}

func init() {
	rootCmd.AddCommand(remCmd)
}
