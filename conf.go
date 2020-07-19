package main

import "time"

var conf Conf

type Conf struct {
	Debug         bool
	Timeout       time.Duration
	Provider      string
	RecordType    string
	Cron          string
	Domain        string
	Host          string
	GoDaddyKey    string
	GoDaddySecret string
	NameComUser   string
	NameComToken  string
	GandiApiKey   string
}
