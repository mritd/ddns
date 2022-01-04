package main

import (
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	ProviderNameCom    = "namecom"
	ProviderGoDaddy    = "godaddy"
	ProviderGandi      = "gandi"
	ProviderCloudflare = "cloudflare"
)

type Conf struct {
	Cron      string        `mapstructure:"cron"`
	Providers ProviderConf  `mapstructure:"provider"`
	Domains   []*DomainConf `mapstructure:"domains"`
}

type DomainConf struct {
	Type     string `mapstructure:"type"`
	Prefix   string `mapstructure:"prefix"`
	Domain   string `mapstructure:"domain"`
	Provider string `mapstructure:"provider"`

	provider Provider
}

type ProviderConf struct {
	Timeout            time.Duration `mapstructure:"timeout"`
	GoDaddyKey         string        `mapstructure:"go_daddy_key"`
	GoDaddySecret      string        `mapstructure:"go_daddy_secret"`
	NameComUser        string        `mapstructure:"name_com_user"`
	NameComToken       string        `mapstructure:"name_com_token"`
	GandiApiKey        string        `mapstructure:"gandi_api_key"`
	CloudflareApiToken string        `mapstructure:"cloudflare_api_token"`
	CloudflareZoneID   string        `mapstructure:"cloudflare_zone_id"`
}

func (conf *Conf) initProvider() {
	for _, d := range conf.Domains {
		var p Provider
		var err error

		switch strings.ToLower(d.Provider) {
		//case ProviderNameCom:
		//	return NewNameCom(c)
		//case ProviderGoDaddy:
		//	return NewGoDaddy(c)
		case ProviderGandi:
			p, err = NewGandi(&conf.Providers)
		//case ProviderCloudflare:

		default:
			logrus.Fatalf("unsupported provider: %s", d.Provider)
		}

		if err != nil {
			logrus.Fatalf("failed to create provider [%s]: %v", d.Provider, err)
		}
		d.provider = p
	}
}

type Record struct {
	Type   string
	Domain string
	Prefix string
	Value  string
}

func (r Record) String() string {
	return r.Prefix + "." + r.Domain
}
