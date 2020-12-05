package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/mritd/logger"

	jsoniter "github.com/json-iterator/go"

	"github.com/robfig/cron"
)

const ApiIpsb = "https://api.ip.sb/ip"

func run(conf *Conf) error {
	logger.Info("ddns running...")
	logger.Debugf("dns provider: %s", conf.Provider)

	provider, err := GetProvider(conf)
	if err != nil {
		return err
	}

	logger.Debugf("request current ip api: %s", ApiIpsb)
	req, _ := http.NewRequest("GET", ApiIpsb, nil)
	client := http.Client{Timeout: conf.Timeout}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	currentIP := strings.TrimSpace(string(bs))
	logger.Infof("current ip: %s", currentIP)

	addr, err := provider.Query()
	if err != nil {
		if _, ok := err.(RecordNotFoundErr); ok {
			logger.Warnf("not found dns record: %s.%s, creating...", conf.Host, conf.Domain)
			if err := provider.Create(currentIP); err != nil {
				return err
			}
			logger.Infof("create dns record: %s.%s success", conf.Host, conf.Domain)
			return nil
		} else {
			return err
		}
	}

	logger.Infof("record ip: %s", addr)
	if addr != currentIP {
		logger.Infof("record changing...")
		if err := provider.Update(currentIP); err != nil {
			return err
		}
		logger.Infof("dns record changed to %s", currentIP)
	} else {
		logger.Infof("skip...")
	}

	return nil
}

func start(conf *Conf) error {
	c := cron.New()
	err := c.AddFunc(conf.Cron, func() {
		if err := run(conf); err != nil {
			logger.Error(err)
		}
	})
	if err != nil {
		return err
	}

	c.Start()
	logger.Info("ddns started.")
	if conf.Debug {
		confJson, _ := jsoniter.MarshalToString(conf)
		logger.Debug(confJson)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	c.Stop()
	logger.Info("ddns exit.")
	return nil
}
