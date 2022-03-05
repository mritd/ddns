package main

import "github.com/sirupsen/logrus"

type Conf struct {
	Cron   string
	ApiKey string
	Type   string
	Domain string
	Prefix []string

	provider Provider
}

func (c *Conf) initProvider() {
	p, err := NewGandi(c)
	if err != nil {
		logrus.Fatalf("failed to create gandi client: %v", err)
	}
	c.provider = p
}

type Record struct {
	Type   string
	Domain string
	Prefix string
	Value  string
}

func (r Record) String() string {
	return r.Prefix + "." + r.Domain
}
