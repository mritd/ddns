package cmd

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

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
	rootCmd.PersistentFlags().StringVarP(&provider, "provider", "p", "namecom", "dns service provider")
	rootCmd.PersistentFlags().StringVarP(&cron, "cron", "c", "@every 5m", "ddns check crontab")
	rootCmd.PersistentFlags().StringVarP(&recordType, "record", "r", "A", "domain record type")
	rootCmd.PersistentFlags().StringVar(&apiKey, "key", "", "dns service provider api key")
	rootCmd.PersistentFlags().StringVar(&apiSecret, "secret", "", "dns service provider api secret")
	rootCmd.PersistentFlags().StringSliceVar(&domain, "domain", nil, "domain A record")
}
