package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

func CheckSHA512(s string) string {
	h := sha512.Sum512([]byte(s))  
	return hex.EncodeToString(h[:])
}
