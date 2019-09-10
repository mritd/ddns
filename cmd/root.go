package cmd

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var debug bool
var timeout time.Duration
var provider string
var recordType string
var cron string
var apiKey, apiSecret string
var domain []string

var rootCmd = &cobra.Command{
	Use:   "ddns",
	Short: "DDNS Tool",
	Long: `
DDNS Tool.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initLog)
	rootCmd.PersistentFlags().StringVarP(&provider, "provider", "p", "namecom", "dns service provider")
	rootCmd.PersistentFlags().StringVarP(&cron, "cron", "c", "@every 5m", "ddns check crontab")
	rootCmd.PersistentFlags().StringVarP(&recordType, "record", "r", "A", "domain record type")
	rootCmd.PersistentFlags().StringVar(&apiKey, "key", "", "dns service provider api key")
	rootCmd.PersistentFlags().StringVar(&apiSecret, "secret", "", "dns service provider api secret")
	rootCmd.PersistentFlags().StringSliceVar(&domain, "domain", nil, "domain records")
	rootCmd.PersistentFlags().DurationVar(&timeout, "timeout", 3*time.Second, "http request timeout")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug mode")
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
