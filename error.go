package main

import "fmt"

type RecordNotFoundErr struct {
	Host   string
	Domain string
}

func (rerr RecordNotFoundErr) Error() string {
	return fmt.Sprintf("records [%s.%s] not found", rerr.Host, rerr.Domain)
}

func NewRecordNotFoundErr(host, domain string) RecordNotFoundErr {
	return RecordNotFoundErr{
		Host:   host,
		Domain: domain,
	}
}

type HttpRequestErr struct {
	Code    int
	Message string
}

func (herr HttpRequestErr) Error() string {
	return fmt.Sprintf("http request failed, code: %d, message: %s", herr.Code, herr.Message)
}

func NewHttpRequestErr(code int, message string) HttpRequestErr {
	return HttpRequestErr{
		Code:    code,
		Message: message,
	}
}
