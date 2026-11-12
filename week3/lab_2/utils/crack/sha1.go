package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

func CheckSHA1(s string) string {
	h := sha1.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}
