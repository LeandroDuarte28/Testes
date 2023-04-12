package services_test

import (
	"testing"
)

func TestSum(t *testing.T) {
	if services.sum(2, 2) != 4 {
		t.Error("Expected 4")
	}
}
