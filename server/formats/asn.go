package formats

import (
	"encoding/asn1"
)

func MarshalASN(data interface{}) ([]byte, error){
	asnMarshalData, err := asn1.Marshal(data)
	return asnMarshalData, err
}

func UnMarshalASN(asnData []byte, structure interface{}) error{
	_, err := asn1.Unmarshal(asnData, structure)
	if err != nil {
		return err
	}
	return nil
}

