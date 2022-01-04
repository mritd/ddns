package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/go-resty/resty/v2"

	"github.com/sirupsen/logrus"

	"github.com/robfig/cron"
)

const ApiIpsb = "https://api.ip.sb/ip"

func run(cli *resty.Client, conf *Conf) {
	logrus.Debugf("request current ip api: %s", ApiIpsb)
	resp, err := cli.R().Get(ApiIpsb)
	if err != nil || resp.IsError() {
		logrus.Errorf("failed to query current ip: %v, %v", err, resp.Error())
		return
	}

	currentIP := resp.String()
	logrus.Infof("current ip: %s", currentIP)

	for _, d := range conf.Domains {
		reqr := &Record{
			Type:   d.Type,
			Domain: d.Domain,
			Prefix: d.Prefix,
		}
		logrus.Infof("checking record: %s", reqr)
		r, err := d.provider.Query(reqr)
		if err != nil {
			if _, ok := err.(RecordNotFoundErr); ok {
				logrus.Warnf("dns record not found: %v, creating...", reqr)
				reqr.Value = currentIP
				if err = d.provider.Create(reqr); err != nil {
					logrus.Error(err)
				}
				continue
			}
		}

		logrus.Infof("record: %s -> %s", reqr, r.Value)
		if r.Value != currentIP {
			logrus.Infof("dns record changing: %s -> %s ", reqr, currentIP)
			reqr.Value = currentIP
			if err = d.provider.Update(reqr); err != nil {
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
