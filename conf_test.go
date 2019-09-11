package main

import (
	"os"
	"time"
)

func initConf() {
	conf = Conf{
		Debug:      true,
		Timeout:    10 * time.Second,
		Provider:   os.Getenv("DDNS_PROVIDER"),
		RecordType: os.Getenv("DDNS_RECORD_TYPE"),
		Cron:       "@every 5m",
		ApiKey:     os.Getenv("DDNS_KEY"),
		ApiSecret:  os.Getenv("DDNS_SECRET"),
		Domain:     os.Getenv("DDNS_DOMAIN"),
		Host:       os.Getenv("DDNS_HOST"),
	}
}
