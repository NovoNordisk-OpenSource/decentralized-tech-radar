/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Fetcher"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch <Url> <Branch> <Whitelist_Filepath> [Url1] [Branch1] [Whitelist_Filepath1]",
	Short: "fetch one or more files from a Git repository",
	Long: `The fetcher is used to pull whitelisted files/folders from one or more git repositories. It takes a string containing 3 values:

	1. A URL to a git based repository
	2. A branch name
	3. A path to a whitelist file
	`,
	
	Args: cobra.MinimumNArgs(3),

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) % 3 != 0 {
			panic("arguments is not divisable by 3")
		} 
		fmt.Println("")
		Fetcher.ListingReposForFetch(args)
		
		fmt.Println("\nFetch complete.")
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
