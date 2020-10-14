// Package encrypt 加密 rsa非对称加密 速度较慢
package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var privateKey, publicKey []byte

// RsaEncrypt 加密
func RsaEncrypt(originData []byte) ([]byte, error) {
	// 解析公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := publicInterface.(*rsa.PublicKey)
	// 加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, originData)
}

// RsaDecrypt 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	// 解析私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// RsaEncryptString Rsa加密后转化为Base64
func RsaEncryptString(originData string) (string, error) {
	cipherArr, err := RsaEncrypt([]byte(originData))
	if err != nil {
		return "", err
	} else {
		return Base64EncodeByte(cipherArr), nil
	}
}

// RsaDecryptString 先Base解码再Rsa解密
func RsaDecryptString(ciphertext string) (string, error) {
	cipherArr := Base64Decode(ciphertext)
	originArr, err := RsaDecrypt([]byte(cipherArr))
	if err != nil {
		return "", err
	} else {
		return string(originArr), nil
	}
}
