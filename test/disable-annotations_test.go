package test

import (
	"testing"

	"github.com/remerge/revive/lint"
	"github.com/remerge/revive/rule"
)

func TestDisabledAnnotations(t *testing.T) {
	testRule(t, "disable-annotations", &rule.ExportedRule{}, &lint.RuleConfig{})
}
