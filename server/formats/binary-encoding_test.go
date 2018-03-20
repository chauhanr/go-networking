package formats

import "testing"

func TestBinaryEncoding(t *testing.T){
	tests := []struct{
		data []byte
	}{
		{[]byte{1,2,3,4,5,6,7,8}},
		{[]byte{1,2,3,4,5,6,7,8,9,10,11,12}},
	}

	for _, tc := range tests{
		encodedBuffer, err := BinaryEncode(tc.data)
		if err != nil{
			t.Fatalf("Error encoding data found error %s", err.Error())
		}
		data, err := BinaryDecode(encodedBuffer, len(tc.data))
		t.Logf("decoded data %v", data)
		for i, ch := range tc.data {
			if ch != data[i]{
				t.Errorf("Expected %d but got %d", ch, data[i])
			}
		}
	}
}