package reader

import (
	"dragon-compiler/golex/spec"
	"errors"
	"os"
	"strings"
)

const sectionDelimiter = "%%"

// ParseFromString 解析 lex 文本并拆分为 definitions/rules/user code 三段。
func ParseFromString(content string) (spec.LexSpec, error) {
	segments := strings.Split(content, sectionDelimiter)
	if len(segments) < 2 {
		return spec.LexSpec{}, errors.New("invalid lex spec: missing section delimiter %%")
	}

	parsedSpec := spec.LexSpec{
		Definitions: strings.TrimSpace(segments[0]),
		Rules:       strings.TrimSpace(segments[1]),
	}

	if len(segments) > 2 {
		parsedSpec.UserCode = strings.TrimSpace(strings.Join(segments[2:], sectionDelimiter))
	}

	return parsedSpec, nil
}

// ParseFromFile 从文件读取 lex 内容并解析。
func ParseFromFile(filePath string) (spec.LexSpec, error) {
	content, readErr := os.ReadFile(filePath)
	if readErr != nil {
		return spec.LexSpec{}, readErr
	}

	return ParseFromString(string(content))
}
