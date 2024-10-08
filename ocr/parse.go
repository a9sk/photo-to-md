package ocr

import (
	"strings"
)

func parseHOCR(hocr string) []StyledText {
	lines := strings.Split(hocr, "\n")
	var styledTexts []StyledText

	for _, line := range lines {
		if strings.Contains(line, "<b>") {
			styledTexts = append(styledTexts, StyledText{
				Text:  extractText(line, "<b>", "</b>"),
				Style: "bold",
			})
		} else if strings.Contains(line, "<i>") {
			styledTexts = append(styledTexts, StyledText{
				Text:  extractText(line, "<i>", "</i>"),
				Style: "italic",
			})
		} else {
			styledTexts = append(styledTexts, StyledText{
				Text:  line,
				Style: "normal",
			})
		}
	}

	return styledTexts
}

func extractText(line string, startTag string, endTag string) string {
	start := strings.Index(line, startTag)
	end := strings.Index(line, endTag)
	if start >= 0 && end > start {
		return line[start+len(startTag) : end]
	}
	return ""
}
