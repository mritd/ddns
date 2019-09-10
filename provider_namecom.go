package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	jsoniter "github.com/json-iterator/go"
)

const NameComApi = "https://api.name.com"

type NameCom struct {
	client *http.Client
	conf   Conf
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
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v4/domains/%s/records", NameComApi, name.conf.Domain), nil)
	if err != nil {
		return NameRecord{}, err
	}

	req.SetBasicAuth(name.conf.ApiKey, name.conf.ApiSecret)

	resp, err := name.client.Do(req)
	if err != nil {
		return NameRecord{}, err
	}
	defer func() { _ = resp.Body.Close() }()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		logrus.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		return NameRecord{}, errors.New(fmt.Sprintf("name.com api request failed, code: %d, message: %s", resp.StatusCode, buf.String()))
	}

	var records []NameRecord
	err = jsoniter.UnmarshalFromString(jsoniter.Get(buf.Bytes(), "records").ToString(), &records)
	if err != nil {
		return NameRecord{}, err
	}

	for _, r := range records {
		if r.Host == name.conf.Host && r.Type == name.conf.RecordType {
			return r, nil
		}
	}
	return NameRecord{}, errors.New(fmt.Sprintf("records [%s.%s] not found", conf.Host, conf.Domain))
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

	payload := fmt.Sprintf(`{"host":"%s","type":"%s","answer":"%s","ttl":300}`, name.conf.Host, name.conf.RecordType, ip)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/v4/domains/%s/records/%d", NameComApi, name.conf.Domain, r.ID), bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	req.SetBasicAuth(name.conf.ApiKey, name.conf.ApiSecret)

	resp, err := name.client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		logrus.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("name.com api request failed, code: %d, message: %s", resp.StatusCode, buf.String()))
	}

	return nil
}

func (name *NameCom) Create(ip string) error {

	payload := fmt.Sprintf(`{"host":"%s","type":"%s","answer":"%s","ttl":300}`, name.conf.Host, name.conf.RecordType, ip)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v4/domains/%s/records", NameComApi, name.conf.Domain), bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	req.SetBasicAuth(name.conf.ApiKey, name.conf.ApiSecret)

	resp, err := name.client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		logrus.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("name.com api request failed, code: %d, message: %s", resp.StatusCode, buf.String()))
	}

	return nil
}

func NewNameCom(conf Conf) *NameCom {
	return &NameCom{
		client: &http.Client{
			Timeout: conf.Timeout,
		},
		conf: conf,
	}
}
