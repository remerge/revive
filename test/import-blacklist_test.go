package test

import (
	"testing"

	"github.com/remerge/revive/lint"
	"github.com/remerge/revive/rule"
)

func TestImportsBlacklist(t *testing.T) {
	args := []interface{}{"crypto/md5", "crypto/sha1"}

	testRule(t, "imports-blacklist", &rule.ImportsBlacklistRule{}, &lint.RuleConfig{
		Arguments: args,
	})
}
