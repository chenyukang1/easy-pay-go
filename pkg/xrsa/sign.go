package xrsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
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
