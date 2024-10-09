package ocr

import (
	"encoding/xml"
	"fmt"
	"photo-to-md/common"
	"strings"
)

func parseHOCR(hocr string) ([]common.StyledText, error) {

	lines := strings.Split(hocr, "\n")

	var styledTexts []common.StyledText

	for _, line := range lines {
		if strings.Contains(line, "<b>") {
			styledTexts = append(styledTexts, common.StyledText{
				Text:  extractText(line, "<b>", "</b>") + "\n",
				Style: "bold",
			})
		} else if strings.Contains(line, "<i>") {
			styledTexts = append(styledTexts, common.StyledText{
				Text:  extractText(line, "<i>", "</i>") + "\n",
				Style: "italic",
			})
		} else if strings.Contains(line, "bbox") {

			bbox := extractBBox(line)

			if bbox.height > 40 {
				styledTexts = append(styledTexts, common.StyledText{
					Text:  extractText(line, "SOL", "EOL") + "\n",
					Style: "title",
				})
			} else if bbox.height > 30 {
				styledTexts = append(styledTexts, common.StyledText{
					Text:  extractText(line, "SOL", "EOL") + "\n",
					Style: "semi-title",
				})
			} else {
				styledTexts = append(styledTexts, common.StyledText{
					Text:  extractText(line, "SOL", "EOL") + "\n",
					Style: "normal",
				})
			}
		} else {
			styledTexts = append(styledTexts, common.StyledText{
				Text:  line + "\n",
				Style: "normal",
			})
		}
	}

	return styledTexts, nil
}

func parseXML(xmlFile string) (string, error) {

	xmlFile = extractXML(xmlFile)

	byteValue := []byte(xmlFile)

	var ocrPage common.OcrPage

	err := xml.Unmarshal(byteValue, &ocrPage)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling XML: %v", err)
	}

	var fullText strings.Builder

	for _, area := range ocrPage.Areas {
		var blockText strings.Builder
		for _, block := range area.Blocks {
			var paragraphText strings.Builder
			for _, line := range block.Lines {
				for _, word := range line.Words {
					paragraphText.WriteString(word.Text + " ")
				}
			}
			blockText.WriteString(paragraphText.String())
		}
		fullText.WriteString(area.Title + " SOL " + blockText.String() + " EOL " + "\n")

	}
	fmt.Print(fullText.String()) //! remove this
	return fullText.String(), nil
}
