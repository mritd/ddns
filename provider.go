package main

import (
	"errors"
	"strings"
)

const (
	ProviderNameCom = "namecom"
	ProviderGoDaddy = "godaddy"
	ProviderGandi   = "gandi"
)

type Provider interface {
	Query() (string, error)
	Update(ip string) error
	Create(ip string) error
}

func GetProvider(conf *Conf) (Provider, error) {
	switch strings.ToLower(conf.Provider) {
	case ProviderNameCom:
		return NewNameCom(conf)
	case ProviderGoDaddy:
		return NewGoDaddy(conf)
	case ProviderGandi:
		return NewGandi(conf)
	default:
		return nil, errors.New("unsupported provider: " + conf.Provider)
	}
}
