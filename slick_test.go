package slick_test

import (
	"testing"

	"github.com/mcandre/slick"
)

func TestVersion(t *testing.T) {
	if slick.Version == "" {
		t.Errorf("Expected slick version to be non-blank")
	}
}
