package main

import (
	"strings"
)

const (
	ProviderNameCom = "namecom"
	ProviderGodaddy = "godaddy"
)

type Provider interface {
	Query() (string, error)
	Update(ip string) error
	Create(ip string) error
}

func GetProvider() Provider {
	switch strings.ToLower(conf.Provider) {
	case ProviderNameCom:
		return NewNameCom()
	case ProviderGodaddy:
		return NewGodaddy()
	default:
		return nil
	}
}
