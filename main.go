package main

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version   string
	buildDate string
	commitID  string
)

var configFile string
var debug bool

var rootCmd = &cobra.Command{
	Use:     "ddns",
	Short:   "Simple DDNS Tool",
	Version: fmt.Sprintf("%s %s %s", version, commitID, buildDate),
	RunE: func(cmd *cobra.Command, args []string) error {
		conf.initProvider()
		return start(&conf)
	},
}

var conf Conf

func init() {
	cobra.OnInitialize(initLog, initConf)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "ddns.yaml", "ddns config file")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug mode")
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

func initConf() {
	viper.SetConfigName("ddns")

	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc")
	viper.AddConfigPath("/etc/ddns")
	_ = viper.ReadInConfig()

	viper.SetEnvPrefix("DDNS")
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&conf); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
