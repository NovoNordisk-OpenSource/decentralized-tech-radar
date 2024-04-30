package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var remCmd = &cobra.Command{
	Use:   "remove <path/to/csvfile> <name> <quadrant>",
	Short: "Removes a line from a csv spec file",
	Long: `Removes the specified tool's line from the given csv specification file. The csv file's header must be in the following format:
name,ring,quadrant,isNew,moved,description`,
	Args: cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		oldFile, err := os.Open(args[0])
		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		scanner := bufio.NewScanner(oldFile)

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			if !strings.Contains(strings.ToLower(line), strings.ToLower(args[1])) &&
				!strings.Contains(strings.ToLower(line), strings.ToLower(args[2])) {
				buf.Write([]byte(line + "\n"))
			}
		}

		err = oldFile.Close()
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(args[0], buf.Bytes(), 0644)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(remCmd)
}
