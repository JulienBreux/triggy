package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	defaultConfigFile = "triggy.yml"
)

var (
	debug      = false
	configFile = defaultConfigFile

	version = "dev"
	commit  = "dev"
	date    = "n/a"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "triggy",
	Short: "Triggy is an awesome tool to connect some actions written in Go.",
	Long:  `Triggy is an awesome tool to connect some actions written in Go.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(buildVersion, buildCommit, buildDate string) {
	version = buildVersion
	commit = buildCommit
	date = buildDate

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug mode")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config-file", "c", defaultConfigFile, "Configuration file")
	rootCmd.Version = version

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
