package test

import (
	"testing"

	"github.com/remerge/revive/rule"
)

// ConstantLogicalExpr rule.
func TestConstantLogicalExpr(t *testing.T) {
	testRule(t, "constant-logical-expr", &rule.ConstantLogicalExprRule{})
}
