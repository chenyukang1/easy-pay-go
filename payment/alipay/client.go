package alipay

import (
	"context"
	"crypto/rsa"
	"easy-pay-go/pkg/xhttp"
	"easy-pay-go/pkg/xlog"
	"easy-pay-go/pkg/xrsa"
	"fmt"
	"net/http"
	"strings"
)

type Config struct {
	AppId              string //appid 是支付宝开放平台应用的唯一标识。 例如：小程序的 appid、网页移动应用的 appid、生活号的 appid、第三方应用的 appid 等。
	AppCertSN          string
	AliPayPublicCertSN string
	AliPayRootCertSN   string
	AppAuthToken       string
	SignType           string          // 商户生成签名字符串所使用的签名算法类型，RSA2
	Charset            string          // 请求使用的编码格式
	isSandbox          bool            // 是否沙箱环境
	privateKey         *rsa.PrivateKey // 商户私钥
	alipayPublicKey    *rsa.PublicKey  // 支付宝公钥
}

type Client[T Response] struct {
	config Config
	hc     *http.Client
	logger *xlog.Logger
}

func NewClient[T Response](config Config) *Client[T] {
	return &Client[T]{
		config: config,
		hc:     xhttp.NewClient(),
		logger: xlog.NewLogger(),
	}
}

func DefaultConfig(appid, privateKey string, isSandbox bool) (*Config, error) {
	key := xrsa.FormatAlipayKey(privateKey, true)
	decode, err := xrsa.DecodePrivateKey([]byte(key))
	if err != nil {
		return nil, err
	}
	cfg := &Config{
		AppId:      appid,
		SignType:   RSA2,
		Charset:    UTF8,
		isSandbox:  isSandbox,
		privateKey: decode,
	}
	return cfg, nil
}

func urlEncode[T Response](cfg *Config, request Request[T]) (string, error) {
	params, err := parsePublicParams(cfg)
	if err != nil {
		return "", err
	}

	if request.GetApiMethodName() == "" {
		return "", &IllegalApiParamError{Msg: "Api Method Name empty"}
	}
	params["method"] = request.GetApiMethodName()
	if request.GetApiVersion() == "" {
		params["version"] = "1"
	} else {
		params["version"] = request.GetApiVersion()
	}

	var builder strings.Builder
	for k, v := range params {
		builder.WriteString(k)
		builder.WriteString("=")
		builder.WriteString(v)
		builder.WriteString("&")
	}
	res := builder.String()
	return res[:len(res)-1], nil
}

func parsePublicParams(cfg *Config) (map[string]string, error) {
	if cfg.AppId == "" {
		return nil, &IllegalApiParamError{Msg: "AppId empty"}
	}
	m := make(map[string]string)
	m["app_id"] = cfg.AppId

	if cfg.Charset == "" {
		m["charset"] = UTF8
	} else {
		m["charset"] = cfg.Charset
	}

	if cfg.SignType == "" {
		m["sign_type"] = RSA2
	} else {
		m["sign_type"] = cfg.SignType
	}

	return m, nil
}

func (c *Client[T]) Execute(ctx context.Context, req Request[T]) (T, error) {
	var zero T
	url, err := urlEncode(&c.config, req)
	if err != nil {
		return zero, err
	}
	url = GatewayUrl
	if c.config.isSandbox {
		url = GatewaySandboxUrl
	}
	fmt.Println(url)

	return zero, nil
}
