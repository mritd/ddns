package main

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ddns",
	Short: "DDNS Tool",
	Long:  "DDNS Tool.",
	Run:   func(cmd *cobra.Command, args []string) { Run() },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initLog)
	rootCmd.Flags().StringVarP(&conf.Provider, "provider", "p", "gandi", "dns service provider")
	rootCmd.Flags().StringVarP(&conf.Cron, "cron", "c", "@every 5m", "ddns check crontab")
	rootCmd.Flags().StringVar(&conf.RecordType, "recordtype", "A", "domain record type")
	rootCmd.Flags().StringVar(&conf.GoDaddyKey, "godaddy-key", "", "godaddy api key")
	rootCmd.Flags().StringVar(&conf.GoDaddySecret, "godaddy-secret", "", "godaddy api secret")
	rootCmd.Flags().StringVar(&conf.NameComUser, "namecom-user", "", "namecom api user name")
	rootCmd.Flags().StringVar(&conf.NameComToken, "namecom-token", "", "namecom api token")
	rootCmd.Flags().StringVar(&conf.GandiApiKey, "gandi-key", "", "gandi api key")
	rootCmd.Flags().StringVar(&conf.Host, "host", "", "domain hosts")
	rootCmd.Flags().StringVar(&conf.Domain, "domain", "", "domain name")
	rootCmd.Flags().DurationVar(&conf.Timeout, "timeout", 10*time.Second, "http request timeout")
	rootCmd.Flags().BoolVar(&conf.Debug, "debug", false, "debug mode")

	_ = rootCmd.MarkFlagRequired("host")
	_ = rootCmd.MarkFlagRequired("domain")
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
