package xrsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"hash"
)

type SignType int

const (
	RSA SignType = iota
	RSA2
)

// SignWithRSA2 商户私钥签名
func SignWithRSA2(privateKey *rsa.PrivateKey, message string) (string, error) {
	h := sha256.Sum256([]byte(message))
	sig, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, h[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}

// VerifyWithRSA2 支付宝公钥验签
func VerifyWithRSA2(pubKey *rsa.PublicKey, message []byte, signatureBase64 string) error {
	sig, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return err
	}
	h := sha256.Sum256(message)
	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, h[:], sig)
}

func Sign(bytes []byte, privateKey *rsa.PrivateKey, signType SignType) (signature string, err error) {
	var (
		h         hash.Hash
		algo      crypto.Hash
		encrypted []byte
	)
	if bytes == nil || privateKey == nil {
		return "", errors.New("bytes or privateKey nil")
	}
	switch signType {
	case RSA:
		h = sha1.New()
		algo = crypto.SHA1
	case RSA2:
		h = sha256.New()
		algo = crypto.SHA256
	default:
		h = sha256.New()
		algo = crypto.SHA256
	}
	if _, err = h.Write(bytes); err != nil {
		return
	}
	if encrypted, err = rsa.SignPKCS1v15(rand.Reader, privateKey, algo, h.Sum(nil)); err != nil {
		return
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}
