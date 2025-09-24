package alipay

import (
	"context"
	"crypto/rsa"
	"easy-pay-go/pkg/xcrypto/xrsa"
	"easy-pay-go/pkg/xhttp"
	"easy-pay-go/pkg/xlog"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"
)

type Config struct {
	AppId              string //appid 是支付宝开放平台应用的唯一标识。 例如：小程序的 appid、网页移动应用的 appid、生活号的 appid、第三方应用的 appid 等。
	AppCertSN          string
	AliPayPublicCertSN string
	AliPayRootCertSN   string
	PrivateKey         *rsa.PrivateKey // 应用私钥
	AlipayPublicKey    *rsa.PublicKey  // 支付宝公钥
	AppAuthToken       string
	SignType           string // 商户生成签名字符串所使用的签名算法类型，RSA2
	Format             string // 仅支持json
	Charset            string // 请求使用的编码格式
	isSandbox          bool   // 是否沙箱环境
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
	if appid == "" {
		return nil, &IllegalApiParamError{Msg: "AppId empty"}
	}
	if privateKey == "" {
		return nil, &IllegalApiParamError{Msg: "PrivateKey empty"}
	}

	key := xrsa.FormatAlipayKey(privateKey, true)
	decode, err := xrsa.DecodePrivateKey([]byte(key))
	if err != nil {
		return nil, err
	}
	cfg := &Config{
		AppId:      appid,
		SignType:   RSA2,
		Charset:    UTF8,
		Format:     JSON,
		isSandbox:  isSandbox,
		PrivateKey: decode,
	}
	return cfg, nil
}

func buildQueryParams[T Response](cfg *Config, request Request[T]) (queryParams string, err error) {
	var params map[string]any
	params, err = buildPubParams(cfg)
	if err != nil {
		return "", err
	}
	if request.GetApiMethodName() == "" {
		return "", &IllegalApiParamError{Msg: "Api Method Name empty"}
	}
	if request.GetApiVersion() == "" {
		return "", &IllegalApiParamError{Msg: "Api Version empty"}
	}

	var bizContent string
	if request.HasBizContent() && request.GetBizModel() != nil {
		if request.NeedEncrypt() {
			// TODO
		} else {
			bizContent, err = request.GetBizModel().ToJson()
			if err != nil {
				return "", err
			}
		}
		params["biz_content"] = bizContent
	}

	params["method"] = request.GetApiMethodName()
	params["version"] = request.GetApiVersion()
	if request.GetReturnUrl() != "" {
		params["return_url"] = request.GetReturnUrl()
	}
	if request.GetNotifyUrl() != "" {
		params["notify_url"] = request.GetNotifyUrl()
	}

	var signature string
	if signature, err = sign(params, cfg.PrivateKey); err != nil {
		return
	}
	params["sign"] = signature

	return urlEncode(params, true), nil
}

func sign(params map[string]any, privateKey *rsa.PrivateKey) (signature string, err error) {
	encode := urlEncode(params, false)

	var signType xrsa.SignType
	switch params["sign_type"] {
	case RSA2:
		signType = xrsa.RSA2
	default:
		signType = xrsa.RSA
	}

	signature, err = xrsa.Sign([]byte(encode), privateKey, signType)
	return
}

func buildPubParams(cfg *Config) (map[string]any, error) {
	if cfg.AppId == "" {
		return nil, &IllegalApiParamError{Msg: "AppId empty"}
	}
	m := make(map[string]any)
	m["app_id"] = cfg.AppId
	m["timestamp"] = time.Now().Format(TimeFormat)

	if cfg.Charset == "" {
		m["charset"] = UTF8
	} else {
		m["charset"] = cfg.Charset
	}
	if cfg.Format == "" {
		m["format"] = JSON
	} else {
		m["format"] = cfg.Format
	}
	if cfg.SignType == "" {
		m["sign_type"] = RSA2
	} else {
		m["sign_type"] = cfg.SignType
	}
	if cfg.AppAuthToken != "" {
		m["app_auth_token"] = cfg.AppAuthToken
	}
	if cfg.AppCertSN != "" {
		m["app_cert_sn"] = cfg.AppCertSN
	}
	if cfg.AliPayRootCertSN != "" {
		m["alipay_root_cert_sn"] = cfg.AliPayRootCertSN
	}
	return m, nil
}

func urlEncode(params map[string]any, encode bool) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var builder strings.Builder
	for _, k := range keys {
		if v := params[k]; v != nil {
			s, ok := v.(string)
			if !ok {
				s = fmt.Sprintf("%v", v)
			}
			if encode {
				k = url.QueryEscape(k)
				s = url.QueryEscape(s)
			}

			builder.WriteString(k)
			builder.WriteString("=")
			builder.WriteString(s)
			builder.WriteString("&")
		}
	}
	res := builder.String()
	return res[:len(res)-1]
}

func (c *Client[T]) PageExecute(ctx context.Context, req Request[T]) (t T, err error) {
	var (
		zero        T
		queryParams string
		builder     strings.Builder
	)
	if queryParams, err = buildQueryParams[T](&c.config, req); err != nil {
		return zero, err
	}

	builder.WriteString("https://")
	path := GatewayPath
	if c.config.isSandbox {
		path = GatewaySandboxPath
	}
	builder.WriteString(path)
	builder.WriteByte('?')
	builder.WriteString(queryParams)

	switch req.GetHttpMethod() {
	case xhttp.HttpGet:
		resp := reflect.New(req.GetResponseType().Elem()).Interface().(T)
		resp.SetBody(builder.String())
		return resp, nil
	}

	return zero, nil
}
