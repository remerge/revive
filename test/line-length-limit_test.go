package test

import (
	"testing"

	"github.com/remerge/revive/lint"
	"github.com/remerge/revive/rule"
)

func TestLineLengthLimit(t *testing.T) {
	testRule(t, "line-length-limit", &rule.LineLengthLimitRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(100)},
	})
}
