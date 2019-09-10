package main

import (
	"os"
	"testing"
	"time"
)

func exampleConfig() Conf {
	return Conf{
		Debug:      false,
		Timeout:    3 * time.Second,
		Provider:   os.Getenv("DDNS_PROVIDER"),
		RecordType: os.Getenv("DDNS_RECORD_TYPE"),
		Cron:       "@every 10s",
		ApiKey:     os.Getenv("DDNS_KEY"),
		ApiSecret:  os.Getenv("DDNS_SECRET"),
		Domain:     os.Getenv("DDNS_DOMAIN"),
		Host:       os.Getenv("DDNS_HOST"),
	}
}

func TestNameCom_Create(t *testing.T) {
	err := NewNameCom(exampleConfig()).Create("1.1.1.1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNameCom_Update(t *testing.T) {
	err := NewNameCom(exampleConfig()).Update("2.2.2.2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNameCom_Query(t *testing.T) {
	r, err := NewNameCom(exampleConfig()).Query()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(r)
	}
}
