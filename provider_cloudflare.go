package main

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

const CloudflareAPI = "https://api.cloudflare.com/client/v4/"

type Cloudflare struct {
	cli *resty.Client
}

func (p *Cloudflare) Query(r *Record) (*Record, error) {
	return nil, errors.New("not implemented")
}

func (p *Cloudflare) Update(*Record) error {
	return errors.New("not implemented")
}

func (p *Cloudflare) Create(*Record) error {
	return errors.New("not implemented")
}

func NewCloudflare(p *ProviderConf) (*Cloudflare, error) {
	if p.GandiApiKey == "" {
		return nil, errors.New("gand api key is empty")
	}

	cli := resty.New().
		SetTimeout(p.Timeout).
		SetAuthScheme("Apikey").
		SetAuthToken(p.GandiApiKey)

	if debug {
		EnableTrace(cli)
	}
	return &Cloudflare{
		cli: cli,
	}, nil
}
