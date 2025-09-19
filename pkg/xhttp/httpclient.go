package xhttp

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

func NewClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second, // TCP连接超时
				KeepAlive: 30 * time.Second, // 空闲连接保持时间
			}).DialContext,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			MaxIdleConnsPerHost:   1000, // 最大空闲连接数
			MaxConnsPerHost:       3000,
			IdleConnTimeout:       90 * time.Second, // 空闲连接超时
			TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时
			ExpectContinueTimeout: 1 * time.Second,
			DisableKeepAlives:     false, // 开启连接复用
			ForceAttemptHTTP2:     true,
		},
	}
}
