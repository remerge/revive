package test

import (
	"testing"

	"github.com/remerge/revive/rule"
)

func TestModifiesValRec(t *testing.T) {
	testRule(t, "modifies-value-receiver", &rule.ModifiesValRecRule{})
}
