package xrsa

import (
	"fmt"
	"strings"
)

const (
	rawLen                = 64
	privateKeyPlaceHolder = "PRIVATE KEY"
	publicKeyPlaceHolder  = "PUBLIC KEY"
)

func FormatAlipayKey(key string, private bool) string {
	var placeholder string
	if private {
		placeholder = privateKeyPlaceHolder
	} else {
		placeholder = publicKeyPlaceHolder
	}
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("-----BEGIN RSA %s-----\n", placeholder))
	lines := len(key) / rawLen
	start, end := 0, rawLen
	for i := 0; i < lines; i++ {
		builder.WriteString(key[start:end])
		builder.WriteByte('\n')
		start += rawLen
		end += rawLen
	}
	if len(key)%rawLen > 0 {
		builder.WriteString(key[start:end])
	}
	builder.WriteString(fmt.Sprintf("-----END RSA %s-----\n", placeholder))
	return builder.String()
}
