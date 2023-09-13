package tests

import (
	"fmt"
	"github.com/goal-web/example/app/usecase"
	"github.com/tj/assert"
	"testing"
)

func TestGenerateRsaKey(t *testing.T) {
	priKey, pubKey, err := usecase.GenerateRSAKey()
	assert.NoError(t, err)
	fmt.Println(string(priKey), string(pubKey))
}
