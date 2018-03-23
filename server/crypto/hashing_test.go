package crypto

import (
	"testing"
)

func TestHashString(t *testing.T) {
	tests := []struct{
		strValue string
		hashScheme string
		hashedValue string
	}{
		{"hello\n", "MD5", "b1946ac9 2492d234 7c6235b4 d2611184 "},

	}

	for _, tc := range tests{
		hashedValue := HashString(tc.hashScheme, tc.strValue)

		if hashedValue != tc.hashedValue {
			t.Errorf("Expected hash value %s but got hashed value %s", tc.hashedValue, hashedValue)
		}
	}

}
