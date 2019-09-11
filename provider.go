package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	ProviderNameCom = "namecom"
	ProviderGodaddy = "godaddy"
	ProviderAliyun  = "aliyun"
)

var RecordNotFoundErr = errors.New(fmt.Sprintf("records not found"))

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
		return NewNameCom()
	case ProviderAliyun:
		return NewNameCom()
	default:
		return nil
	}
}
