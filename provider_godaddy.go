package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

const GoDaddyAPI = "https://api.godaddy.com"

type GoDaddy struct {
	client *http.Client
}

type GoDaddyRecord struct {
	Data     string `json:"data"`
	Name     string `json:"name"`
	TTL      int    `json:"ttl"`
	Type     string `json:"type"`
	Priority int    `json:"priority,omitempty"`
}

func (p *GoDaddy) query() (GoDaddyRecord, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/domains/%s/records", GoDaddyAPI, conf.Domain), nil)
	if err != nil {
		return GoDaddyRecord{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("sso-key %s:%s", conf.GoDaddyKey, conf.GoDaddySecret))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	resp, err := p.client.Do(req)
	if err != nil {
		return GoDaddyRecord{}, err
	}
	defer func() { _ = resp.Body.Close() }()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GoDaddyRecord{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return GoDaddyRecord{}, fmt.Errorf("request failed, status code: %d, message: %s", resp.StatusCode, string(bs))
	}

	var records []GoDaddyRecord
	err = jsoniter.Unmarshal(bs, &records)
	if err != nil {
		return GoDaddyRecord{}, err
	}

	for _, r := range records {
		if r.Name == conf.Host && r.Type == conf.RecordType {
			return r, nil
		}
	}
	return GoDaddyRecord{}, NewRecordNotFoundErr(conf.Host, conf.Domain)

}

func (p *GoDaddy) Query() (string, error) {
	r, err := p.query()
	if err != nil {
		return "", err
	} else {
		return r.Data, nil
	}
}

func (p *GoDaddy) Update(ip string) error {
	payload := fmt.Sprintf(`[{"data":"%s","ttl":600}]`, ip)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/v1/domains/%s/records/%s/%s", GoDaddyAPI, conf.Domain, conf.RecordType, conf.Host), bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("sso-key %s:%s", conf.GoDaddyKey, conf.GoDaddySecret))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		bs, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("request failed, status code: %d, message: %s", resp.StatusCode, string(bs))
	}

	return nil
}

func (p *GoDaddy) Create(ip string) error {
	payload := fmt.Sprintf(`[{"data":"%s","name":"%s","ttl":600,"type":"%s"}]`, ip, conf.Host, conf.RecordType)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v1/domains/%s/records", GoDaddyAPI, conf.Domain), bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("sso-key %s:%s", conf.GoDaddyKey, conf.GoDaddySecret))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		bs, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("request failed, status code: %d, message: %s", resp.StatusCode, string(bs))
	}

	return nil
}

func NewGoDaddy() (*GoDaddy, error) {
	if conf.GoDaddyKey == "" || conf.GoDaddySecret == "" {
		return nil, errors.New("godaddy api key or api secret is empty")
	}
	return &GoDaddy{
		client: &http.Client{
			Timeout: conf.Timeout,
		},
	}, nil
}
