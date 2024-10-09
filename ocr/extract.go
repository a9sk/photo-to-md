package ocr

import (
	"fmt"
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

type BBox struct {
	width  int
	height int
}

func extractBBox(line string) BBox {
	var x1, y1, x2, y2 int
	if _, err := fmt.Sscanf(line, `bbox %d %d %d %d`, &x1, &y1, &x2, &y2); err != nil {
		return BBox{0, 0}
	}
	return BBox{
		width:  x2 - x1,
		height: y2 - y1, //normal 25, usage 32, title 40
	}
}

func extractXML(hocrHTML string) string {
	var sb strings.Builder
	sb.WriteString(extractText(hocrHTML, "<body>", "</body>"))
	return sb.String()
}
