package usecase

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GenerateRSAKey() ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

	//将数据保存到文件
	privateBytes := pem.EncodeToMemory(&pem.Block{Type: "privateKey", Bytes: x509PrivateKey})

	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	publicBytes := pem.EncodeToMemory(&pem.Block{Type: "publicKey", Bytes: X509PublicKey})

	return privateBytes, publicBytes, nil
}
