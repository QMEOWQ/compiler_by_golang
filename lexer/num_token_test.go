package lexer

import (
	"testing"
)

func TestNumToken(t *testing.T) {
	num_token := NewNumToken(123)
	if num_token.value != 123 {
		t.Fatalf("expected num token value 123, got %d", num_token.value)
	}
	if num_token.ToString() != "123" {
		t.Fatalf("expected num token string 123, got %s", num_token.ToString())
	}
	num_tag := num_token.Tag
	if num_tag.ToString() != "NUM" {
		t.Fatalf("expected num token tag NUM, got %s", num_tag.ToString())
	}
}

func TestRealToken(t *testing.T) {
	real_token := NewRealToken(3.1415926)
	if real_token.value != 3.1415926 {
		t.Fatalf("expected real token value 3.1415926, got %f", real_token.value)
	}
	if real_token.ToString() != "3.1415926" {
		t.Fatalf("expected real token string 3.1415926, got %s", real_token.ToString())
	}

	real_tag := real_token.Tag
	if real_tag.ToString() != "REAL" {
		t.Fatalf("expected real token tag REAL, got %s", real_tag.ToString())
	}
}
