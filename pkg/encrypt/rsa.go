package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func RsaEncrypt(pubPEM, msg string) ([]byte, error) {
	//pem解码
	block, _ := pem.Decode([]byte(pubPEM))

	input := block.Bytes

	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(input)
	if err!=nil{
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(msg))
	if err!=nil{
		return nil, err
	}

	return cipherText, nil
}