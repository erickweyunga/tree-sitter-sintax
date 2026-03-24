package tree_sitter_sintax_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-sintax"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_sintax.Language())
	if language == nil {
		t.Errorf("Error loading Sintax grammar")
	}
}
