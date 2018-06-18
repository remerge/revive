package test

import (
	"testing"

	"github.com/remerge/revive/lint"
	"github.com/remerge/revive/rule"
)

func TestArgumentLimit(t *testing.T) {
	testRule(t, "argument-limit", &rule.ArgumentsLimitRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(3)},
	})
}
