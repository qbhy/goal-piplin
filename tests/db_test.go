package tests

import (
	"fmt"
	"github.com/qbhy/goal-piplin/app/models"
	"testing"
)

func TestQuery(t *testing.T) {
	initApp()
	fmt.Print(models.Users().Count())
}
