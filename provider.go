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

func GetProvider() (Provider, error) {
	switch strings.ToLower(conf.Provider) {
	case ProviderNameCom:
		return NewNameCom()
	case ProviderGoDaddy:
		return NewGoDaddy()
	case ProviderGandi:
		return NewGandi()
	default:
		return nil, errors.New("unsupported provider: " + conf.Provider)
	}
}
