package http

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

var Client = &http.Client{
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   time.Second * 5,
			KeepAlive: time.Second * 30,
		}).DialContext,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		IdleConnTimeout:     time.Second * 30,
		ForceAttemptHTTP2:   true,
		TLSHandshakeTimeout: time.Second * 10,
	},
	Timeout: time.Second * 15,
}
