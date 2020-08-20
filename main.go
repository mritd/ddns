package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/mritd/zaplogger"

	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

var (
	version   string
	buildDate string
	commitID  string

	versionTpl = `
Name: ddns
Version: %s
Arch: %s
BuildDate: %s
CommitID: %s
`
)

var logger *zap.SugaredLogger

var rootCmd = &cobra.Command{
	Use:     "ddns",
	Version: version,
	Short:   "DDNS Tool",
	Run:     func(cmd *cobra.Command, args []string) { Run() },
}

func main() {
	// init zap logger
	zc, err := zaplogger.NewConfig(zapConf)
	if err != nil {
		log.Fatalf("Failed to create zap logger config: %v", err)
	}
	logger = zaplogger.NewLogger(zc).Sugar()

	// run command
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
	}
}

func init() {
	// zap logger
	rootCmd.PersistentFlags().BoolVar(&zapConf.Development, "zap-devel", false, "Enable zap development mode (changes defaults to console encoder, debug log level, disables sampling and stacktrace from 'warning' level)")
	rootCmd.PersistentFlags().StringVar(&zapConf.Encoder, "zap-encoder", "console", "Zap log encoding ('json' or 'console')")
	rootCmd.PersistentFlags().StringVar(&zapConf.Level, "zap-level", "info", "Zap log level (one of 'debug', 'info', 'warn', 'error')")
	rootCmd.PersistentFlags().BoolVar(&zapConf.Sample, "zap-sample", false, "Enable zap log sampling. Sampling will be disabled for log level is debug")
	rootCmd.PersistentFlags().StringVar(&zapConf.TimeEncoding, "zap-time-encoding", "default", "Sets the zap time format ('default', 'epoch', 'millis', 'nano', or 'iso8601')")
	rootCmd.PersistentFlags().StringVar(&zapConf.StackLevel, "zap-stacktrace-level", "error", "Set the minimum log level that triggers stacktrace generation")

	// ddns provider config
	rootCmd.Flags().StringVarP(&conf.Provider, "provider", "p", "gandi", "dns service provider")
	rootCmd.Flags().StringVar(&conf.GoDaddyKey, "godaddy-key", "", "godaddy api key")
	rootCmd.Flags().StringVar(&conf.GoDaddySecret, "godaddy-secret", "", "godaddy api secret")
	rootCmd.Flags().StringVar(&conf.NameComUser, "namecom-user", "", "namecom api user name")
	rootCmd.Flags().StringVar(&conf.NameComToken, "namecom-token", "", "namecom api token")
	rootCmd.Flags().StringVar(&conf.GandiApiKey, "gandi-key", "", "gandi api key")

	// ddns config
	rootCmd.Flags().StringVarP(&conf.Cron, "cron", "c", "@every 5m", "ddns check crontab")
	rootCmd.Flags().StringVar(&conf.RecordType, "recordtype", "A", "domain record type")
	rootCmd.Flags().StringVar(&conf.Host, "host", "", "domain hosts")
	rootCmd.Flags().StringVar(&conf.Domain, "domain", "", "domain name")
	rootCmd.Flags().DurationVar(&conf.Timeout, "timeout", 10*time.Second, "http request timeout")
	rootCmd.Flags().BoolVar(&conf.Debug, "debug", false, "debug mode")

	// version template
	rootCmd.SetVersionTemplate(fmt.Sprintf(versionTpl, version, runtime.GOOS+"/"+runtime.GOARCH, buildDate, commitID))

	_ = rootCmd.MarkFlagRequired("host")
	_ = rootCmd.MarkFlagRequired("domain")
}
