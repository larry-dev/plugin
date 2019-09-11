package encrypt

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
	"time"
)

func TestGenSha1(t *testing.T) {
	appKey := "admin"
	nonce := uuid.NewV4().String()
	curTime := fmt.Sprintf("%v", time.Now().Unix())
	t.Log(appKey, nonce, curTime, GetSha1(appKey+nonce+curTime))
}

func TestGenPassword(t *testing.T) {

	t.Log(GetSha1("admin"))
}
