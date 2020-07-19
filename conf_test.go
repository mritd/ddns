package main

import (
	"os"
	"time"
)

func initConf() {
	conf = Conf{
		Debug:         true,
		Cron:          "@every 5m",
		Timeout:       10 * time.Second,
		Host:          os.Getenv("DDNS_HOST"),
		Domain:        os.Getenv("DDNS_DOMAIN"),
		Provider:      os.Getenv("DDNS_PROVIDER"),
		RecordType:    os.Getenv("DDNS_RECORD_TYPE"),
		GoDaddyKey:    os.Getenv("DDNS_GODADDY_KEY"),
		GoDaddySecret: os.Getenv("DDNS_DODADDY_SECRET"),
		NameComUser:   os.Getenv("DDNS_NAMECOM_USER"),
		NameComToken:  os.Getenv("DDNS_NAMECOM_TOKEN"),
		GandiApiKey:   os.Getenv("DDNS_GANDI_KEY"),
	}
}
