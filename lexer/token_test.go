package lexer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testTokenName(t *testing.T) {
	real_token := NewToken(REAL)
	require.Equal(t, "REAL", real_token.ToString())

	id := NewToken(ID)
	require.Equal(t, "indetifier", id.ToString())
}
