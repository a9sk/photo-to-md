package ocr

import (
	"encoding/xml"
	"fmt"
	"photo-to-md/common"
	"strings"
)

// function to parse hocr and return styled text. it extracts different text styles (bold, italic, titles) based on the hocr markup.
func parseHOCR(hocr string) ([]common.StyledText, error) {

	// parse the hocr string using xml parsing function.
	var err error
	hocr, err = parseXML(hocr)
	// handle any errors during parsing.
	if err != nil {
		// return an empty styled text with error.
		var emptyRet []common.StyledText
		emptyRet = append(emptyRet, common.StyledText{
			Text:  "",
			Style: "",
		})
		return emptyRet, fmt.Errorf("%v", err)
	}

	// split the hocr content by lines.
	lines := strings.Split(hocr, "\n")

	// create a slice to store the styled text output.
	var styledTexts []common.StyledText

	// iterate over each line and check for text styles (bold, italic, bbox).
	for _, line := range lines {
		// following checks are useless as Tesseract (for how it is configured)
		// does not recognize bold/italic text most of the time.
		// // check if the line contains bold text.
		// if strings.Contains(line, "<b>") {
		// 	// extract the bold text and append it to styledTexts.
		// 	styledTexts = append(styledTexts, common.StyledText{
		// 		Text:  extractText(line, "<b>", "</b>") + "\n",
		// 		Style: "bold",
		// 	})
		// 	// check if the line contains italic text.
		// } else if strings.Contains(line, "<i>") {
		// 	// extract the italic text and append it to styledTexts.
		// 	styledTexts = append(styledTexts, common.StyledText{
		// 		Text:  extractText(line, "<i>", "</i>") + "\n",
		// 		Style: "italic",
		// 	})
		// 	// check if the line contains bbox information.
		// } else
		if strings.Contains(line, "bbox") {

			// extract the bbox data from the line.
			bbox := extractBBox(line)

			// determine the style based on the bbox height.
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
			// append normal text (lines without special styles).
			styledTexts = append(styledTexts, common.StyledText{
				Text:  line + "\n",
				Style: "normal",
			})
		}
	}

	// return the final styled text output.
	return styledTexts, nil
}

// function to parse xml content from a string. returns a formatted string with text and styling.
func parseXML(xmlFile string) (string, error) {

	// extract relevant xml data from the input file.
	xmlFile = extractXML(xmlFile)

	// convert the xml string to byte array.
	byteValue := []byte(xmlFile)

	// define a struct to hold the unmarshalled xml content.
	var ocrPage common.OcrPage

	// unmarshal the xml content into the struct.
	err := xml.Unmarshal(byteValue, &ocrPage)
	// handle any errors during unmarshaling.
	if err != nil {
		return "", fmt.Errorf("error unmarshaling XML: %v", err)
	}

	// use a string builder to efficiently construct the full text.
	var fullText strings.Builder

	// iterate over the areas in the parsed ocr data.
	for _, area := range ocrPage.Areas {
		var blockText strings.Builder
		for _, block := range area.Blocks {
			var paragraphText strings.Builder
			// iterate over the lines within blocks.
			for _, line := range block.Lines {
				var lineText strings.Builder
				// extract words from each line.
				for _, word := range line.Words {
					lineText.WriteString(word.Text + " ")
				}
				// format lines to undeline:
				// the bbox sizes (which show the hight of the text on the single line)
				// the Start Of Line (SOL) and the End Of Line (EOL) to easily work with them
				paragraphText.WriteString(line.Title + " SOL " + lineText.String() + " EOL " + "\n")
			}
			blockText.WriteString(paragraphText.String())
		}
		fullText.WriteString(blockText.String())
	}
	// return the full text as a string.
	return fullText.String(), nil
}
