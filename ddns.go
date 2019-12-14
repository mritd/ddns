package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	jsoniter "github.com/json-iterator/go"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

const IPSBAPI = "https://api.ip.sb/ip"

func run() error {

	logrus.Info("ddns running...")
	logrus.Debugf("dns provider: %s", conf.Provider)

	provider := GetProvider()
	if provider == nil {
		return errors.New(fmt.Sprintf("failed to get provider: %s", conf.Provider))
	}

	logrus.Debugf("request current ip api: %s", IPSBAPI)
	req, _ := http.NewRequest("GET", IPSBAPI, nil)
	client := http.Client{Timeout: conf.Timeout}
	resp, err := client.Do(req)
	if err != nil {
		return NewHttpRequestErr(-1, err.Error())
	}
	defer func() { _ = resp.Body.Close() }()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return err
	}

	currentIP := strings.TrimSpace(buf.String())
	logrus.Infof("current ip: %s", currentIP)

	addr, err := provider.Query()
	if err != nil {
		if _, ok := err.(RecordNotFoundErr); ok {
			logrus.Warnf("not found dns record: %s.%s, creating...", conf.Host, conf.Domain)
			err = provider.Create(currentIP)
			if err != nil {
				return err
			}
			logrus.Infof("create dns record: %s.%s success", conf.Host, conf.Domain)
			return nil
		} else {
			return err
		}
	}

	logrus.Infof("record ip: %s", addr)

	if addr != currentIP {
		logrus.Infof("changing...")
		err = provider.Update(currentIP)
		if err != nil {
			return err
		}
		logrus.Infof("dns record changed to %s", currentIP)
	} else {
		logrus.Infof("skip...")
	}

	return nil
}

func Run() {

	c := cron.New()
	err := c.AddFunc(conf.Cron, func() {
		err := run()
		if err != nil {
			logrus.Error(err)
		}
	})
	if err != nil {
		logrus.Fatal(err)
	}

	c.Start()
	logrus.Info("ddns started.")
	if conf.Debug {
		confJson, _ := jsoniter.MarshalToString(conf)
		logrus.Debug(confJson)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	c.Stop()
	logrus.Info("ddns exit.")
}
