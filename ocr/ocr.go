package ocr

import (
	"fmt"
	"os"
	"os/exec"
	"photo-to-md/common"
	"strings"
)

func PerformOCR(imagePath string) ([]common.StyledText, error) {
	fmt.Println("Using Tesseract to retrive text from the image")

	isSpecific := os.Getenv("ISSPECIFIC")

	var styledOutput []common.StyledText

	if isSpecific == "true" {
		var err error
		styledOutput, err = dohOCR(imagePath)
		if err != nil {
			var emptyRet []common.StyledText
			emptyRet = append(emptyRet, common.StyledText{
				Text:  "",
				Style: "",
			})
			return emptyRet, fmt.Errorf("error running the hOCR: %v", err)
		}
	} else {
		output, err := doOCR(imagePath)
		if err != nil {
			var emptyRet []common.StyledText
			emptyRet = append(emptyRet, common.StyledText{
				Text:  "",
				Style: "",
			})
			return emptyRet, fmt.Errorf("error running the OCR: %v", err)
		}
		styledOutput = append(styledOutput, common.StyledText{
			Text:  output,
			Style: "",
		})
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

func dohOCR(imagePath string) ([]common.StyledText, error) {
	fmt.Println("Performing hOCR on image...")

	cmd := exec.Command("tesseract", imagePath, "stdout", "hocr")

	output, err := cmd.Output()
	if err != nil {

		var emptyRet []common.StyledText
		emptyRet = append(emptyRet, common.StyledText{
			Text:  "",
			Style: "",
		})

		return emptyRet, fmt.Errorf("error running Tesseract: %v", err)
	}

	parsedOutput, err := parseHOCR(string(output))
	if err != nil {
		var emptyRet []common.StyledText
		emptyRet = append(emptyRet, common.StyledText{
			Text:  "",
			Style: "",
		})
		return emptyRet, fmt.Errorf("%v", err)
	}

	return parsedOutput, nil
}
