package test

import (
	"testing"

	"github.com/remerge/revive/rule"
)

func TestDeepExit(t *testing.T) {
	testRule(t, "deep-exit", &rule.DeepExitRule{})
}
