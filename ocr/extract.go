package ocr

import (
	"strings"
)

func extractText(line string, startTag string, endTag string) string {
	start := strings.Index(line, startTag)
	end := strings.Index(line, endTag)
	if start >= 0 && end > start {
		return line[start+len(startTag) : end]
	}
	return ""
}
