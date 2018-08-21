package test

import (
	"testing"

	"github.com/remerge/revive/rule"
)

// TestModifiesParam rule.
func TestModifiesParam(t *testing.T) {
	testRule(t, "modifies-param", &rule.ModifiesParamRule{})
}
