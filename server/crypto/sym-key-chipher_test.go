package crypto

import "testing"

func TestSymetricKeyChiper (t *testing.T){
	symKey := "chauhan"
	message := "ritesh\n\n"

	enc, err := EncryptSymetricKey(symKey, message)
	if err != nil{
		t.Fatalf("Error Encryting message %s", err.Error())
	}

	returnMessage, err := DecryptSymetricKey(symKey, enc)

	if err != nil{
		t.Fatalf("Error Decryting message %s", err.Error())
	}

	if returnMessage != message {
		t.Errorf("Expected message %s but got %s instead", message, returnMessage)
	}
}
