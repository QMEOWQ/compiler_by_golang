package lexer

import (
	"testing"
)

func TestWordToken(t *testing.T) {
	word := NewWordToken("variable", ID)
	if word.ToString() != "variable" {
		t.Fatalf("expected word token string variable, got %s", word.ToString())
	}
	word_tag := word.Tag
	if word_tag.ToString() != "ID" {
		t.Fatalf("expected word token tag ID, got %s", word_tag.ToString())
	}
}

func TestKeyWords(t *testing.T) {
	key_words := GetKeyWords()
	if len(key_words) != 15 {
		t.Fatalf("expected keyword length 15, got %d", len(key_words))
	}

	and_key_word := key_words[0]
	if and_key_word.ToString() != "&&" {
		t.Fatalf("expected first keyword &&, got %s", and_key_word.ToString())
	}

	or_key_word := key_words[1]
	if or_key_word.ToString() != "||" {
		t.Fatalf("expected second keyword ||, got %s", or_key_word.ToString())
	}
}
