package test

import (
	"testing"

	"github.com/remerge/revive/rule"
)

// TestConfusingNaming rule.
func TestConfusingNaming(t *testing.T) {
	testRule(t, "confusing-naming1", &rule.ConfusingNamingRule{})
}
