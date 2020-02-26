package command

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienbreux/triggy/internal/config"
	"github.com/julienbreux/triggy/internal/manager"
	"github.com/julienbreux/triggy/internal/triggy/read"
	"github.com/julienbreux/triggy/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "dev"
	date    = "n/a"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "triggy",
	Short: "Triggy is an awesome tool to connect some actions written in Go.",
	Long:  `Triggy is an awesome tool to connect some actions written in Go.`,
	Run:   rootRun,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(buildVersion, buildCommit, buildDate string) {
	version = buildVersion
	commit = buildCommit
	date = buildDate

	rootCmd.Version = version

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func rootRun(cmd *cobra.Command, args []string) {
	subLog := log.With().Str("process", "global").Logger()
	// Init configuration
	c, err := config.New()
	if err != nil {
		subLog.Error().Err(err).Msg("Unable to load environment variable")
		os.Exit(1)
	}
	if c.LoggerFormat == "text" {
		subLog = subLog.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
	// Init logger
	zerolog.SetGlobalLevel(logger.ParseLevel(c.LoggerLevel))
	if c.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		subLog.Debug().Msg("Debug mode enabled")
	}

	subLog.Debug().RawJSON("config", c.JSON()).Msg("Environment variables loaded")

	// Load configuration file
	t, err := read.Triggy(c.ConfigPath, c.ConfigFile)
	if err != nil {
		subLog.Fatal().Err(err).Msgf("Unable to read %s file on path %s", c.ConfigFile, c.ConfigPath)
	}
	subLog.Debug().RawJSON("config", t.JSON()).Msg("Configuration file loaded")

	// Start Triggy :)
	subLog.Info().Msg("Starting Triggy...")
	mngr := manager.New(t, c)
	mngr.Start()
	subLog.Info().Msg("Triggy started")

	// Stop Triggy :(
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	subLog.Info().Str("reason", fmt.Sprintf("%s", <-ch)).Msg("Stopping Triggy...")

	mngr.Stop()
	subLog.Info().Msg("Triggy stopped")
}
