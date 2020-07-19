package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

const NameComApi = "https://api.name.com"

type NameCom struct {
	client *http.Client
}

type NameRecord struct {
	ID         int    `json:"id"`
	DomainName string `json:"domainName"`
	Host       string `json:"host"`
	Fqdn       string `json:"fqdn"`
	Type       string `json:"type"`
	Answer     string `json:"answer"`
	TTL        int    `json:"ttl"`
}

func (p *NameCom) query() (NameRecord, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v4/domains/%s/records", NameComApi, conf.Domain), nil)
	if err != nil {
		return NameRecord{}, err
	}

	req.SetBasicAuth(conf.NameComUser, conf.NameComToken)

	resp, err := p.client.Do(req)
	if err != nil {
		return NameRecord{}, err
	}
	defer func() { _ = resp.Body.Close() }()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return NameRecord{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return NameRecord{}, fmt.Errorf("request failed, status code: %d, message: %s", resp.StatusCode, string(bs))
	}

	var records []NameRecord
	err = jsoniter.UnmarshalFromString(jsoniter.Get(bs, "records").ToString(), &records)
	if err != nil {
		return NameRecord{}, err
	}

	for _, r := range records {
		if r.Host == conf.Host && r.Type == conf.RecordType {
			return r, nil
		}
	}
	return NameRecord{}, NewRecordNotFoundErr(conf.Host, conf.Domain)
}

func (p *NameCom) Query() (string, error) {
	r, err := p.query()
	if err != nil {
		return "", err
	} else {
		return r.Answer, nil
	}

}

func (p *NameCom) Update(ip string) error {
	r, err := p.query()
	if err != nil {
		return err
	}

	payload := fmt.Sprintf(`{"host":"%s","type":"%s","answer":"%s","ttl":300}`, conf.Host, conf.RecordType, ip)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/v4/domains/%s/records/%d", NameComApi, conf.Domain, r.ID), bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	req.SetBasicAuth(conf.NameComUser, conf.NameComToken)

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

func (p *NameCom) Create(ip string) error {
	payload := fmt.Sprintf(`{"host":"%s","type":"%s","answer":"%s","ttl":300}`, conf.Host, conf.RecordType, ip)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v4/domains/%s/records", NameComApi, conf.Domain), bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	req.SetBasicAuth(conf.NameComUser, conf.NameComToken)

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

func NewNameCom() (*NameCom, error) {
	if conf.NameComUser == "" || conf.NameComToken == "" {
		return nil, errors.New("namecom api user or token is empty")
	}
	return &NameCom{
		client: &http.Client{
			Timeout: conf.Timeout,
		},
	}, nil
}
