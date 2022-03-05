package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

// GandiAPI https://api.gandi.net/docs/livedns/#v5-livedns-domains-fqdn-records-rrset_name-rrset_type
const GandiAPI = "https://api.gandi.net/v5/livedns/domains/%s/records/%s/%s"

type Gandi struct {
	cli *resty.Client
}

type GandiRecord struct {
	RRSetHref   string   `json:"rrset_href"`
	RRSetName   string   `json:"rrset_name"`
	RRSetType   string   `json:"rrset_type"`
	RRSetTTL    int      `json:"rrset_ttl"`
	RRSetValues []string `json:"rrset_values"`
}

func (p *Gandi) Query(r *Record) (*Record, error) {
	var gdr GandiRecord
	resp, err := p.cli.R().SetResult(&gdr).Get(fmt.Sprintf(GandiAPI, r.Domain, r.Prefix, r.Type))
	if err != nil || resp.IsError() {
		if resp.StatusCode() == 404 {
			return nil, NewRecordNotFoundErr(r)
		}

		return nil, fmt.Errorf("failed to query record(%d): %w, %v", resp.StatusCode(), err, resp.Error())
	}

	return &Record{
		Type:   r.Type,
		Domain: r.Domain,
		Prefix: r.Prefix,
		Value:  gdr.RRSetValues[0],
	}, nil
}

func (p *Gandi) Create(r *Record) error {
	payload := fmt.Sprintf(`{"rrset_values":["%s"],"rrset_ttl":300}`, r.Value)
	resp, err := p.cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(fmt.Sprintf(GandiAPI, r.Domain, r.Prefix, r.Type))
	if err != nil || resp.IsError() {
		return fmt.Errorf("faile to update record [%s]: %w, %v", r, err, resp.Error())
	}

	return nil
}

func (p *Gandi) Update(r *Record) error {
	payload := fmt.Sprintf(`{"rrset_values":["%s"],"rrset_ttl":300}`, r.Value)
	resp, err := p.cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Put(fmt.Sprintf(GandiAPI, r.Domain, r.Prefix, r.Type))
	if err != nil || resp.IsError() {
		return fmt.Errorf("faile to create record [%s]: %w, %v", r, err, resp.Error())
	}
	return nil
}

func NewGandi(c *Conf) (*Gandi, error) {
	if c.ApiKey == "" {
		return nil, errors.New("gand api key is empty")
	}

	cli := resty.New().
		SetTimeout(3 * time.Second).
		SetAuthScheme("Apikey").
		SetAuthToken(c.ApiKey)

	if debug {
		EnableTrace(cli)
	}
	return &Gandi{
		cli: cli,
	}, nil
}
