package crypto

import (
	"crypto/md5"
	"hash"
	"golang.org/x/crypto/md4"
	"fmt"
	"bytes"
)

func HashString(scheme string, content string) string{
	bContent := []byte(content)
	var hashScheme hash.Hash

	switch (scheme) {
		case "MD5" :
			    // hash value from md5 algo is a 16byte array and this is printed out in ASCII form in
			    // 4 hexa decimal numbers
				hashScheme = md5.New()
				return md5ProcessingScheme(hashScheme,bContent)
		case "MD4" :
				hashScheme = md4.New()
				return ""
		default :
			hashScheme = md5.New()
			return md5ProcessingScheme(hashScheme,bContent)
	}
}

func md5ProcessingScheme(md5Scheme hash.Hash, bContent []byte) string {
	md5Scheme.Write(bContent)
	hashValue := md5Scheme.Sum(nil)
	hashSize := md5Scheme.Size()

	var hashReturn = ""
	for n :=0; n<hashSize; n +=4 {
		var val uint32
		val = uint32(hashValue[n])<<24+
			uint32(hashValue[n+1])<<16+
			uint32(hashValue[n+2])<<8+
			uint32(hashValue[n+3])
		buf := bytes.NewBufferString("")
		fmt.Fprintf(buf, "%x", val)

		hashReturn = hashReturn + buf.String()+" "
	}
	return hashReturn
}