package lexer

import (
	"testing"
)

func TestTokenName(t *testing.T) {
	indexToken := NewToken(INDEX)
	if indexToken.ToString() != "INDEX" {
		t.Fatalf("expected token string INDEX, got %s", indexToken.ToString())
	}

	real_token := NewToken(REAL)
	if real_token.ToString() != "REAL" {
		t.Fatalf("expected token string REAL, got %s", real_token.ToString())
	}
}
