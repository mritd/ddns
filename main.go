package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mritd/logger"

	"github.com/urfave/cli/v2"
)

var (
	version   string
	buildDate string
	commitID  string
)

func main() {
	app := &cli.App{
		Name:    "ddns",
		Usage:   "DDNS Tool",
		Version: fmt.Sprintf("%s %s %s", version, buildDate, commitID),
		Authors: []*cli.Author{
			{
				Name:  "mritd",
				Email: "mritd@linux.com",
			},
		},
		Copyright: "Copyright (c) 2020 mritd, All rights reserved.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "provider",
				Aliases: []string{"p"},
				Value:   "gandi",
				Usage:   "dns service provider",
				EnvVars: []string{"DDNS_PROVIDER"},
			},
			&cli.StringFlag{
				Name:    "cron",
				Aliases: []string{"c"},
				Value:   "@every 5m",
				Usage:   "ddns check crontab",
				EnvVars: []string{"DDNS_CRON"},
			},
			&cli.StringFlag{
				Name:    "record-type",
				Value:   "A",
				Usage:   "domain record type",
				EnvVars: []string{"DDNS_RECORD_TYPE"},
			},
			&cli.StringFlag{
				Name:     "host",
				Usage:    "domain host",
				Required: true,
				EnvVars:  []string{"DDNS_HOST"},
			},
			&cli.StringFlag{
				Name:    "domain",
				Usage:   "domain name",
				EnvVars: []string{"DDNS_DOMAIN"},
			},
			&cli.DurationFlag{
				Name:    "timeout",
				Usage:   "http request timeout",
				EnvVars: []string{"DDNS_TIMEOUT"},
				Value:   10 * time.Second,
			},
			&cli.StringFlag{
				Name:    "godaddy-key",
				Usage:   "godaddy api key",
				EnvVars: []string{"DDNS_GODADDY_KEY"},
			},
			&cli.StringFlag{
				Name:    "godaddy-secret",
				Usage:   "godaddy api secret",
				EnvVars: []string{"DDNS_GODADDY_SECRET"},
			},
			&cli.StringFlag{
				Name:    "namecom-user",
				Usage:   "namecom api user name",
				EnvVars: []string{"DDNS_NAMECOM_USER"},
			},
			&cli.StringFlag{
				Name:    "namecom-token",
				Usage:   "namecom api token",
				EnvVars: []string{"DDNS_NAMECOM_TOKEN"},
			},
			&cli.StringFlag{
				Name:    "gandi-key",
				Usage:   "gandi api key",
				EnvVars: []string{"DDNS_GANDI_KEY"},
			},
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "debug mode",
				EnvVars: []string{"DDNS_DEBUG"},
				Value:   false,
			},
		},
		Action: func(c *cli.Context) error {
			return start(&Conf{
				Debug:         c.Bool("debug"),
				Timeout:       c.Duration("timeout"),
				Provider:      c.String("provider"),
				RecordType:    c.String("record-type"),
				Cron:          c.String("cron"),
				Domain:        c.String("domain"),
				Host:          c.String("host"),
				GoDaddyKey:    c.String("godaddy-key"),
				GoDaddySecret: c.String("godaddy-secret"),
				NameComUser:   c.String("namecom-user"),
				NameComToken:  c.String("namecom-token"),
				GandiApiKey:   c.String("gandi-key"),
			})
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Error(err)
	}
}
