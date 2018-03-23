package crypto

import (
	"golang.org/x/crypto/blowfish"
	"bytes"
)

func EncryptSymetricKey(key string, content string) ([512]byte,error){
	symKey := []byte(key)
	cipher, err := blowfish.NewCipher(symKey)

	if err != nil{
		return [512]byte{}, err
	}

	bContent := []byte(content)
	var enc [512]byte
	cipher.Encrypt(enc[0:],bContent)

	return enc, nil
}

func DecryptSymetricKey(key string, enc [512]byte) (string, error) {
	symKey := []byte(key)
	var decrypt [8]byte
	cipher, err := blowfish.NewCipher(symKey)
	if err != nil{
		return "", err
	}
	cipher.Decrypt(decrypt[0:], enc[0:])
	result := bytes.NewBuffer(nil)
	result.Write(decrypt[0:8])
	return string(result.Bytes()),nil
}
