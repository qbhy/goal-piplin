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

func TestName(t *testing.T) {
	fmt.Println(len("c9cfe439-4384-4b32-9e2b-56257b916f59"))
}
