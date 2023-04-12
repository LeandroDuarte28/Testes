package services_test

import (
	"testing"

	"github.com/Testes/services"
)

func TestSum(t *testing.T) {
	if (services.Sum(2, 2)) != 4 {
		t.Error("Expected 4")
	}
}
