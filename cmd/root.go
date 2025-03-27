/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-webscraper",
	Short: "A CLI tool to scrape the web + more",
	Long: `This is a simple CLI tool to scrape the web and compare products, among other things...
	
	Additiional functions are:
	- Password generation: 'generate'
	- Scrape the Musiciansfriend website for product info: 'scrape'
	- Trace IP addresses: 'trace'
	- More to come!
	
	Made by:
	 _____  ____    ____  ____   __  _                    
	|     ||    \  /    T|    \ |  l/ ]                   
	|   __j|  D  )Y  o  ||  _  Y|  ' /                    
	|  l_  |    / |     ||  |  ||    \                    
	|   _] |    \ |  _  ||  |  ||     Y                   
	|  T   |  .  Y|  |  ||  |  ||  .  |                   
	l__j   l__j\_jl__j__jl__j__jl__j\_j                   
                                                              
	 _____  ____    ____  ____      __  ____   ___   ____     ___ 
	|     ||    \  /    T|    \    /  ]l    j /   \ |    \   /  _]
	|   __j|  D  )Y  o  ||  _  Y  /  /  |  T Y     Y|  _  Y /  [_ 
	|  l_  |    / |     ||  |  | /  /   |  | |  O  ||  |  |Y    _]
	|   _] |    \ |  _  ||  |  |/   \_  |  | |     ||  |  ||   [_ 
	|  T   |  .  Y|  |  ||  |  |\     | j  l l     !|  |  ||     T
	l__j   l__j\_jl__j__jl__j__j \____j|____j \___/ l__j__jl_____j`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-webscraper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
