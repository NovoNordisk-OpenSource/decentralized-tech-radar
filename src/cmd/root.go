package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tech_radar",
	Short: "A tool for generating tech radar from csv specfiles",
	Long: ` 
	________   ________  ___      ___ ________                                                            
	|\   ___  \|\   __  \|\  \    /  /|\   __  \                                                           
	\ \  \\ \  \ \  \|\  \ \  \  /  / | \  \|\  \                                                          
	 \ \  \\ \  \ \  \\\  \ \  \/  / / \ \  \\\  \                                                         
	  \ \  \\ \  \ \  \\\  \ \    / /   \ \  \\\  \                                                        
	   \ \__\\ \__\ \_______\ \__/ /     \ \_______\                                                       
	    \|__| \|__|\|_______|\|__|/       \|_______|                                                       																							   
	 _________  _______   ________  ___  ___          ________  ________  ________  ________  ________     
	|\___   ___\\  ___ \ |\   ____\|\  \|\  \        |\   __  \|\   __  \|\   ___ \|\   __  \|\   __  \    
	\|___ \  \_\ \   __/|\ \  \___|\ \  \\\  \       \ \  \|\  \ \  \|\  \ \  \_|\ \ \  \|\  \ \  \|\  \   
	     \ \  \ \ \  \_|/_\ \  \    \ \   __  \       \ \   _  _\ \   __  \ \  \ \\ \ \   __  \ \   _  _\  
	      \ \  \ \ \  \_|\ \ \  \____\ \  \ \  \       \ \  \\  \\ \  \ \  \ \  \_\\ \ \  \ \  \ \  \\  \| 
               \ \__\ \ \_______\ \_______\ \__\ \__\       \ \__\\ _\\ \__\ \__\ \_______\ \__\ \__\ \__\\ _\ 
                \|__|  \|_______|\|_______|\|__|\|__|        \|__|\|__|\|__|\|__|\|_______|\|__|\|__|\|__|\|__|

A tool for generating tech radars and manipulating csv files (by fetching and merging).`,
}



// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}


