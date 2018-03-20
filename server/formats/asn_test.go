package formats

import (
	"testing"
)

func TestMarshalUnmarshal(t *testing.T){
	test := []struct{
		data Data
	}{
		{Data{"Ritesh", 24}},
		{Data{"Nitin", 31}},
	}

	for _, tc := range test{
		marshalData, err := MarshalASN(tc.data)
		if err != nil{
			t.Errorf("Data %q should be successfully marshalled but gave error %s", marshalData, err.Error())
		}
		var data = new(Data)
		err = UnMarshalASN(marshalData, data)
		if err != nil{
			t.Fatalf("Data should be unmarshalled but got error %s", err.Error())
		}
		t.Logf("Unmarshalled data is %+v", data)
	}
}
