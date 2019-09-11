package encrypt

import (
	"crypto/sha1"
	"encoding/hex"
)

func GetSha1(v string) string {
	bs := sha1.Sum([]byte(v))
	return hex.EncodeToString(bs[:])
}
