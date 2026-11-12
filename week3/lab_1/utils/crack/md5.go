package crack

import (
	"crypto/md5"
	"encoding/hex"
)

func CheckMD5(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}
