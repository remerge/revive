package test

import (
	"testing"

	"github.com/remerge/revive/lint"
	"github.com/remerge/revive/rule"
)

func TestMaxPublicStructs(t *testing.T) {
	testRule(t, "max-public-structs", &rule.MaxPublicStructsRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(1)},
	})
}
