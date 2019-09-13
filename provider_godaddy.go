package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

const GodaddyAPI = "https://api.godaddy.com"

type Godaddy struct {
	client *http.Client
}

type GodaddyRecord struct {
	Data     string `json:"data"`
	Name     string `json:"name"`
	TTL      int    `json:"ttl"`
	Type     string `json:"type"`
	Priority int    `json:"priority,omitempty"`
}

func (godaddy *Godaddy) query() (GodaddyRecord, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/domains/%s/records", GodaddyAPI, conf.Domain), nil)
	if err != nil {
		return GodaddyRecord{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("sso-key %s:%s", conf.ApiKey, conf.ApiSecret))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	resp, err := godaddy.client.Do(req)
	if err != nil {
		return GodaddyRecord{}, NewHttpRequestErr(-1, err.Error())
	}
	defer func() { _ = resp.Body.Close() }()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		logrus.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		return GodaddyRecord{}, NewHttpRequestErr(resp.StatusCode, buf.String())
	}

	var records []GodaddyRecord
	err = jsoniter.Unmarshal(buf.Bytes(), &records)
	if err != nil {
		return GodaddyRecord{}, err
	}

	for _, r := range records {
		if r.Name == conf.Host && r.Type == conf.RecordType {
			return r, nil
		}
	}
	return GodaddyRecord{}, NewRecordNotFoundErr(conf.Host, conf.Domain)

}

func (godaddy *Godaddy) Query() (string, error) {
	r, err := godaddy.query()
	if err != nil {
		return "", err
	} else {
		return r.Data, nil
	}
}

func (godaddy *Godaddy) Update(ip string) error {

	payload := fmt.Sprintf(`[{"data":"%s","ttl":600}]`, ip)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/v1/domains/%s/records/%s/%s", GodaddyAPI, conf.Domain, conf.RecordType, conf.Host), bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("sso-key %s:%s", conf.ApiKey, conf.ApiSecret))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	resp, err := godaddy.client.Do(req)
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

func (godaddy *Godaddy) Create(ip string) error {

	payload := fmt.Sprintf(`[{"data":"%s","name":"%s","ttl":600,"type":"%s"}]`, ip, conf.Host, conf.RecordType)

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v1/domains/%s/records", GodaddyAPI, conf.Domain), bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("sso-key %s:%s", conf.ApiKey, conf.ApiSecret))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	resp, err := godaddy.client.Do(req)
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

func NewGodaddy() *Godaddy {
	return &Godaddy{
		client: &http.Client{
			Timeout: conf.Timeout,
		},
	}
}
