package ocr

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type StyledText struct {
	Text  string
	Style string
}

func PerformOCR(imagePath string) ([]StyledText, error) {
	fmt.Println("Using Tesseract to retrive text from the image")

	isSpecific := os.Getenv("ISSPECIFIC")

	var styledOutput []StyledText

	if isSpecific == "true" {
		output, err := doOCR(imagePath)
		if err != nil {
			var emptyRet []StyledText
			emptyRet = append(emptyRet, StyledText{
				Text:  "",
				Style: "",
			})
			return emptyRet, fmt.Errorf("error running the OCR: %v", err)
		}
		styledOutput = append(styledOutput, StyledText{
			Text:  output,
			Style: "",
		})

	} else {
		var err error
		styledOutput, err = dohOCR(imagePath)
		if err != nil {
			var emptyRet []StyledText
			emptyRet = append(emptyRet, StyledText{
				Text:  "",
				Style: "",
			})
			return emptyRet, fmt.Errorf("error running the hOCR: %v", err)
		}
	}

	return styledOutput, nil
}

func doOCR(imagePath string) (string, error) {
	fmt.Println("Performing OCR on image...")

	cmd := exec.Command("tesseract", imagePath, "stdout")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running Tesseract: %v", err)
	}

	return strings.TrimSpace(string(output)), nil
}

func dohOCR(imagePath string) ([]StyledText, error) {
	fmt.Println("Performing hOCR on image...")

	cmd := exec.Command("tesseract", imagePath, "stdout", "--hocr")
	output, err := cmd.Output()
	if err != nil {

		var emptyRet []StyledText
		emptyRet = append(emptyRet, StyledText{
			Text:  "",
			Style: "",
		})

		return emptyRet, fmt.Errorf("error running Tesseract: %v", err)
	}

	return parseHOCR(string(output)), nil
}
