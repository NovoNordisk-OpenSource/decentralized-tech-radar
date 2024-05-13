package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <path/to/csvfile> <name> <ring> <quadrant> <isNew> <moved> <description>",
	Short: "adds a new line to a csv file",
	Long: `This command appends a new line to the specified csv file. The csv file's header must be in the following format:
	name,ring,quadrant,isNew,moved,description
	`,
	Args: cobra.MinimumNArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		// Reads data from the file
		old_data, err := os.ReadFile(args[0])
		if err != nil {
			panic(err)
		}

		// Removes the last newline character
		old_data = bytes.Trim(old_data, "\n")

		// Appends the new data to the file
		data := append(old_data, []byte("\n"+strings.Join(args[1:], ","))...)

		// Writes the new data to the file
		err = os.WriteFile(args[0], data, 0644)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Data added: %s", strings.Join(args[1:], ","))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
