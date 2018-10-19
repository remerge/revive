package test

import (
	"testing"

	"github.com/remerge/revive/lint"
	"github.com/remerge/revive/rule"
)

func TestVarNaming(t *testing.T) {
	testRule(t, "var-naming", &rule.VarNamingRule{}, &lint.RuleConfig{
		Arguments: []interface{}{[]interface{}{"ID"}, []interface{}{"VM"}},
	})
}
