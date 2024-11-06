package lexer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testWordToken(t *testing.T) {
	word := NewWordToken("variable", ID)
	require.Equal(t, "variable", word.ToString())
	word_tag := word.Tag
	require.Equal(t, word_tag.ToString(), "identifier")
}

func testWords(t *testing.T) {
	key_words := GetKeyWords()

	and_key_word := key_words[0]
	require.Equal(t, and_key_word.ToString(), "&&")

	or_key_word := key_words[1]
	require.Equal(t, or_key_word.ToString(), "||")
}
