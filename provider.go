package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type Provider interface {
	Query(*Record) (*Record, error)
	Update(*Record) error
	Create(*Record) error
}

func EnableTrace(cli *resty.Client) {
	cli.EnableTrace().OnAfterResponse(
		func(client *resty.Client, resp *resty.Response) error {
			// Explore response object
			logrus.Debug("Response Info:")
			logrus.Debug("  Error      :", resp.Error())
			logrus.Debug("  Status Code:", resp.StatusCode())
			logrus.Debug("  Status     :", resp.Status())
			logrus.Debug("  Proto      :", resp.Proto())
			logrus.Debug("  Time       :", resp.Time())
			logrus.Debug("  Received At:", resp.ReceivedAt())
			logrus.Debug("  Body       :", resp)
			logrus.Debug()

			// Explore trace info
			logrus.Debug("Request Trace Info:")
			ti := resp.Request.TraceInfo()
			logrus.Debug("  URL           :", resp.Request.URL)
			logrus.Debug("  DNSLookup     :", ti.DNSLookup)
			logrus.Debug("  ConnTime      :", ti.ConnTime)
			logrus.Debug("  TCPConnTime   :", ti.TCPConnTime)
			logrus.Debug("  TLSHandshake  :", ti.TLSHandshake)
			logrus.Debug("  ServerTime    :", ti.ServerTime)
			logrus.Debug("  ResponseTime  :", ti.ResponseTime)
			logrus.Debug("  TotalTime     :", ti.TotalTime)
			logrus.Debug("  IsConnReused  :", ti.IsConnReused)
			logrus.Debug("  IsConnWasIdle :", ti.IsConnWasIdle)
			logrus.Debug("  ConnIdleTime  :", ti.ConnIdleTime)
			logrus.Debug("  RequestAttempt:", ti.RequestAttempt)
			logrus.Debug("  RemoteAddr    :", ti.RemoteAddr.String())
			logrus.Debug("  RequestHeader :", resp.Request.Header)

			return nil
		})
}
