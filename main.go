package main

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ddns",
	Short: "DDNS Tool",
	Long: `
DDNS Tool.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initLog)
	rootCmd.PersistentFlags().StringVarP(&conf.Provider, "provider", "p", "namecom", "dns service provider")
	rootCmd.PersistentFlags().StringVarP(&conf.Cron, "cron", "c", "@every 5m", "ddns check crontab")
	rootCmd.PersistentFlags().StringVarP(&conf.RecordType, "record", "r", "A", "domain record type")
	rootCmd.PersistentFlags().StringVar(&conf.ApiKey, "key", "", "dns service provider api key")
	rootCmd.PersistentFlags().StringVar(&conf.ApiSecret, "secret", "", "dns service provider api secret")
	rootCmd.PersistentFlags().StringVar(&conf.Domain, "domain", "", "domain name")
	rootCmd.PersistentFlags().StringVar(&conf.Host, "host", "", "domain hosts")
	rootCmd.PersistentFlags().DurationVar(&conf.Timeout, "timeout", 3*time.Second, "http request timeout")
	rootCmd.PersistentFlags().DurationVarP(&conf.Interval, "interval", "i", 3*time.Minute, "ddns check interval")
	rootCmd.PersistentFlags().BoolVar(&conf.Debug, "debug", false, "debug mode")
}

func initLog() {
	if conf.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
