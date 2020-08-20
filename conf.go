package main

import (
	"time"

	"github.com/mritd/zaplogger"
)

var conf Conf
var zapConf zaplogger.ZapConfig

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
