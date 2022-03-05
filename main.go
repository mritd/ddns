package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var commitID string

var debug bool

var rootCmd = &cobra.Command{
	Use:     "ddns",
	Short:   "Simple DDNS Tool",
	Version: commitID,
	RunE: func(cmd *cobra.Command, args []string) error {
		conf.initProvider()
		return start(&conf)
	},
}

var conf Conf

func init() {
	cobra.OnInitialize(initLog)

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug mode")
	rootCmd.PersistentFlags().StringVarP(&conf.Cron, "cron", "c", "@every 30s", "ddns crontab")
	rootCmd.PersistentFlags().StringVarP(&conf.ApiKey, "key", "k", "", "gandi api key")
	rootCmd.PersistentFlags().StringVarP(&conf.Type, "type", "t", "A", "record type")
	rootCmd.PersistentFlags().StringVarP(&conf.Domain, "domain", "d", "", "base domain")
	rootCmd.PersistentFlags().StringSliceVarP(&conf.Prefix, "prefix", "p", []string{}, "domain prefix")

	_ = rootCmd.PersistentFlags().MarkHidden("type")
}

func initLog() {
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
