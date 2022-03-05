package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/go-resty/resty/v2"

	"github.com/sirupsen/logrus"

	"github.com/robfig/cron"
)

const IPV4API = "https://api-ipv4.ip.sb/ip "

func run(cli *resty.Client, conf *Conf) {
	logrus.Debugf("request current ip api: %s", IPV4API)
	resp, err := cli.R().Get(IPV4API)
	if err != nil || resp.IsError() {
		logrus.Errorf("failed to query current ip: %v, %v", err, resp.Error())
		return
	}

	currentIP := resp.String()
	logrus.Infof("current ip: %s", currentIP)

	for _, p := range conf.Prefix {
		reqr := &Record{
			Type:   conf.Type,
			Domain: conf.Domain,
			Prefix: p,
		}
		logrus.Infof("checking record: %s", reqr)
		r, err := conf.provider.Query(reqr)
		if err != nil {
			if _, ok := err.(RecordNotFoundErr); ok {
				logrus.Warnf("dns record not found: %v, creating...", reqr)
				reqr.Value = currentIP
				if err = conf.provider.Create(reqr); err != nil {
					logrus.Error(err)
				}
				continue
			}
		}

		logrus.Infof("record: %s -> %s", reqr, r.Value)
		if r.Value != currentIP {
			logrus.Infof("dns record changing: %s -> %s ", reqr, currentIP)
			reqr.Value = currentIP
			if err = conf.provider.Update(reqr); err != nil {
				logrus.Error(err)
				continue
			}
			logrus.Infof("dns record changed: %s -> %s", reqr, currentIP)
		} else {
			logrus.Infof("record skiped: %s...", reqr)
		}
	}

}

func start(conf *Conf) error {
	c := cron.New()
	cli := resty.New()
	if debug {
		EnableTrace(cli)
	}

	if err := c.AddFunc(conf.Cron, func() { run(cli, conf) }); err != nil {
		return err
	}

	c.Start()
	logrus.Info("ddns started...")

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	<-ctx.Done()
	c.Stop()
	logrus.Info("ddns exit.")
	return nil
}
