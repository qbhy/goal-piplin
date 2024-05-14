package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"golang.org/x/crypto/ssh"
)

// GenerateRSAKeys 生成2048位的RSA密钥对，并以OpenSSH格式返回公钥和PEM格式返回私钥
func GenerateRSAKeys() (privateKeyPEM string, publicKeySSH string, err error) {
	// 生成密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", "", err
	}

	// 将私钥转换为ASN.1 DER编码
	privateKeyDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// 将DER编码的私钥转换为PEM格式
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyDER,
	}
	privateKeyPEM = string(pem.EncodeToMemory(privateKeyBlock))

	// 从私钥中提取公钥并转换为OpenSSH格式
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	publicKeySSH = string(ssh.MarshalAuthorizedKey(publicKey))

	return
}
