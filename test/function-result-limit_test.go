package test

import (
	"testing"

	"github.com/remerge/revive/lint"
	"github.com/remerge/revive/rule"
)

func TestFunctionResultsLimit(t *testing.T) {
	testRule(t, "function-result-limit", &rule.FunctionResultsLimitRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(3)},
	})
}
