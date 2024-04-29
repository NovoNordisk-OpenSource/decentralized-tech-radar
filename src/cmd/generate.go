package cmd

import (
	html "github.com/NovoNordisk-OpenSource/decentralized-tech-radar/HTML"
	Reader "github.com/NovoNordisk-OpenSource/decentralized-tech-radar/SpecReader"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate <FilePath>",
	Short: "Generates an HTML file that is populated with CSV data",
	Long: `This command reads from a given CSV file, and then generates an HTML file populated with the CSV data.
The FilePath refers to the designated path. An example would be: 'C://Program/MyCSVFile.csv'`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		specs := Reader.CsvToString(args[0])
		html.GenerateHtml(specs)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
