package tests

import (
	"fmt"
	"github.com/goal-web/example/app/models"
	"testing"
)

func TestQuery(t *testing.T) {
	initApp()
	fmt.Print(models.Users().Count())
}
