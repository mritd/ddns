package main

import "time"

var conf Conf

type Conf struct {
	Debug      bool
	Interval   time.Duration
	Timeout    time.Duration
	Provider   string
	RecordType string
	Cron       string
	ApiKey     string
	ApiSecret  string
	Domain     string
	Host       string
}
