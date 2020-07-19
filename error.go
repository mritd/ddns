package main

import "fmt"

type RecordNotFoundErr struct {
	Host   string
	Domain string
}

func (e RecordNotFoundErr) Error() string {
	return fmt.Sprintf("record [%s.%s] not found", e.Host, e.Domain)
}

func NewRecordNotFoundErr(host, domain string) RecordNotFoundErr {
	return RecordNotFoundErr{
		Host:   host,
		Domain: domain,
	}
}
