package formats

import (
	"bytes"
	"encoding/base64"
)

func BinaryEncode(data []byte) (*bytes.Buffer, error){

	encodedBytes := &bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding,encodedBytes)
	_, err := encoder.Write(data)
	err = encoder.Close()
	if err != nil {
		return encodedBytes, err
	}
	return encodedBytes, err
}

func BinaryDecode(encodedData *bytes.Buffer, length int ) ([]byte, error){
	dbuf := make([]byte, length+1)
	decoder := base64.NewDecoder(base64.StdEncoding,encodedData)
	_, err := decoder.Read(dbuf)

	if err != nil{
		return  []byte{}, err
	}
	return dbuf, err
}
