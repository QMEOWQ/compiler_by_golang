package reader

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFromString(t *testing.T) {
	content := `%{
int line;
%}
%%
[0-9]+    { return NUMBER; }
[a-zA-Z]+ { return ID; }
%%
int yywrap() { return 1; }`

	lexSpec, parseErr := ParseFromString(content)
	if parseErr != nil {
		t.Fatalf("expected parse success, got error: %v", parseErr)
	}

	if lexSpec.Definitions == "" {
		t.Fatalf("expected non-empty definitions")
	}
	if lexSpec.Rules == "" {
		t.Fatalf("expected non-empty rules")
	}
	if lexSpec.UserCode == "" {
		t.Fatalf("expected non-empty user code")
	}
}

func TestParseFromStringMissingDelimiter(t *testing.T) {
	content := "no section delimiter here"
	_, parseErr := ParseFromString(content)
	if parseErr == nil {
		t.Fatalf("expected parse failure when delimiter is missing")
	}
}

func TestParseFromFile(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "sample.lex")
	content := `%{
int value;
%}
%%
[a-z]+ { return WORD; }
%%`

	writeErr := os.WriteFile(filePath, []byte(content), 0644)
	if writeErr != nil {
		t.Fatalf("write temp lex file failed: %v", writeErr)
	}

	lexSpec, parseErr := ParseFromFile(filePath)
	if parseErr != nil {
		t.Fatalf("expected file parse success, got error: %v", parseErr)
	}

	if lexSpec.Definitions == "" || lexSpec.Rules == "" {
		t.Fatalf("expected definitions and rules to be non-empty")
	}
}
