package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

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

func (name *NameCom) query() (NameRecord, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v4/domains/%s/records", NameComApi, conf.Domain), nil)
	if err != nil {
		return NameRecord{}, err
	}

	req.SetBasicAuth(conf.ApiKey, conf.ApiSecret)

	resp, err := name.client.Do(req)
	if err != nil {
		return NameRecord{}, NewHttpRequestErr(-1, err.Error())
	}
	defer func() { _ = resp.Body.Close() }()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		logrus.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		return NameRecord{}, NewHttpRequestErr(resp.StatusCode, buf.String())
	}

	var records []NameRecord
	err = jsoniter.UnmarshalFromString(jsoniter.Get(buf.Bytes(), "records").ToString(), &records)
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

func (name *NameCom) Query() (string, error) {
	r, err := name.query()
	if err != nil {
		return "", err
	} else {
		return r.Answer, nil
	}

}

func (name *NameCom) Update(ip string) error {

	r, err := name.query()
	if err != nil {
		return err
	}

	payload := fmt.Sprintf(`{"host":"%s","type":"%s","answer":"%s","ttl":300}`, conf.Host, conf.RecordType, ip)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/v4/domains/%s/records/%d", NameComApi, conf.Domain, r.ID), bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	req.SetBasicAuth(conf.ApiKey, conf.ApiSecret)

	resp, err := name.client.Do(req)
	if err != nil {
		return NewHttpRequestErr(-1, err.Error())
	}
	defer func() { _ = resp.Body.Close() }()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		logrus.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		return NewRecordNotFoundErr(conf.Host, conf.Domain)
	}

	return nil
}

func (name *NameCom) Create(ip string) error {

	payload := fmt.Sprintf(`{"host":"%s","type":"%s","answer":"%s","ttl":300}`, conf.Host, conf.RecordType, ip)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v4/domains/%s/records", NameComApi, conf.Domain), bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	req.SetBasicAuth(conf.ApiKey, conf.ApiSecret)

	resp, err := name.client.Do(req)
	if err != nil {
		return NewHttpRequestErr(-1, err.Error())
	}
	defer func() { _ = resp.Body.Close() }()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		logrus.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		return NewHttpRequestErr(resp.StatusCode, buf.String())
	}

	return nil
}

func NewNameCom() *NameCom {
	return &NameCom{
		client: &http.Client{
			Timeout: conf.Timeout,
		},
	}
}
