package test

import (
	"testing"

	"github.com/remerge/revive/lint"
	"github.com/remerge/revive/rule"
)

func TestRangeValInClosure(t *testing.T) {
	testRule(t, "range-val-in-closure", &rule.RangeValInClosureRule{}, &lint.RuleConfig{})
}
