package encrypt

import (
	"crypto/md5"
	"fmt"
)

func Md5(buf string) string {
	hash := md5.New()
	hash.Write([]byte(buf))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
