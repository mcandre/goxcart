package goxcart_test

import (
	"testing"

	"github.com/mcandre/goxcart"
)

func TestVersion(t *testing.T) {
	if goxcart.Version == "" {
		t.Errorf("Expected version to be non-blank")
	}
}
