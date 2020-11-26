package cmd

import (
	"github.com/krakowski/webwatcher/util"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const (
	version = "1.0.0"
)

var (
	configPath string
)

var rootCommand = &cobra.Command{
	Use:           "webwatcher",
	Short:         "Watches for a list of keywords on a specified website",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := util.ReadConfig(configPath)
		if err != nil {
			log.Fatal(err)
		}

		watcher, err := util.NewWatcher(nil, config)
		if err != nil {
			log.Fatal(err)
		}

		if err := watcher.Watch(); err != nil {
			log.Fatal(err)
		}

		log.Println("Bye!")
	},
}

func init() {
	rootCommand.Version = version
	rootCommand.Flags().StringVar(&configPath, "config", "./config.yaml", "The configuration file to use")
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
