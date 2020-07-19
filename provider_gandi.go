package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

const GandiApi = "https://api.gandi.net"

type Gandi struct {
	client *http.Client
}

type GandiRecord struct {
	RRSetHref   string   `json:"rrset_href"`
	RRSetName   string   `json:"rrset_name"`
	RRSetType   string   `json:"rrset_type"`
	RRSetTTL    int      `json:"rrset_ttl"`
	RRSetValues []string `json:"rrset_values"`
}

func (p *Gandi) query() (GandiRecord, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v5/livedns/domains/%s/records/%s", GandiApi, conf.Domain, conf.Host), nil)
	if err != nil {
		return GandiRecord{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", conf.GandiApiKey))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	resp, err := p.client.Do(req)
	if err != nil {
		return GandiRecord{}, err
	}
	defer func() { _ = resp.Body.Close() }()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GandiRecord{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return GandiRecord{}, fmt.Errorf("request failed, status code: %d, message: %s", resp.StatusCode, string(bs))
	}

	var records []GandiRecord
	err = jsoniter.Unmarshal(bs, &records)
	if err != nil {
		return GandiRecord{}, err
	}

	for _, r := range records {
		if r.RRSetName == conf.Host {
			if len(r.RRSetValues) == 0 {
				return GandiRecord{}, NewRecordNotFoundErr(conf.Host, conf.Domain)
			}
			return r, nil
		}
	}

	return GandiRecord{}, NewRecordNotFoundErr(conf.Host, conf.Domain)
}

func (p *Gandi) Query() (string, error) {
	r, err := p.query()
	if err != nil {
		return "", err
	} else {
		return r.RRSetValues[0], nil
	}
}

func (p *Gandi) Update(ip string) error {
	payload := fmt.Sprintf(`{"items":[{"rrset_type":"%s","rrset_values":["%s"],"rrset_ttl":300}]}`, conf.RecordType, ip)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/v5/livedns/domains/%s/records/%s", GandiApi, conf.Domain, conf.Host), bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", conf.GandiApiKey))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bs, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("request failed, status code: %d, message: %s", resp.StatusCode, string(bs))
	}

	return nil
}

func (p *Gandi) Create(ip string) error {
	payload := fmt.Sprintf(`{"rrset_type":"%s","rrset_values":["%s"],"rrset_ttl":300}`, conf.RecordType, ip)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v5/livedns/domains/%s/records/%s", GandiApi, conf.Domain, conf.Host), bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", conf.GandiApiKey))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bs, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("request failed, status code: %d, message: %s", resp.StatusCode, string(bs))
	}

	return nil
}

func NewGandi() (*Gandi, error) {
	if conf.GandiApiKey == "" {
		return nil, errors.New("gand api key is empty")
	}
	return &Gandi{
		client: &http.Client{
			Timeout: conf.Timeout,
		},
	}, nil
}
