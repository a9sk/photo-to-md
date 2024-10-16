package ocr

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/a9sk/photo-to-md/common"
)

// function to perform ocr on an image. returns styled text and an error.
func PerformOCR(imagePath string) ([]common.StyledText, error) {
	// print message indicating tesseract is being used.
	fmt.Println("Using Tesseract to retrieve text from the image")

	// get environment variable to check if specific ocr should be used.
	isSpecific := os.Getenv("ISSPECIFIC")

	// declare an empty slice of styled text for the output.
	var styledOutput []common.StyledText

	// check if the specific ocr mode is enabled.
	if isSpecific == "true" {
		// perform hocr if specific mode is enabled.
		var err error
		styledOutput, err = dohOCR(imagePath)
		// handle errors in running hocr.
		if err != nil {
			// return an empty styled text with error.
			var emptyRet []common.StyledText
			emptyRet = append(emptyRet, common.StyledText{
				Text:  "",
				Style: "",
			})
			return emptyRet, fmt.Errorf("error running the hOCR: %v", err)
		}
	} else {
		// perform regular ocr if specific mode is not enabled.
		output, err := doOCR(imagePath)
		// handle errors in running ocr.
		if err != nil {
			// return an empty styled text with error.
			var emptyRet []common.StyledText
			emptyRet = append(emptyRet, common.StyledText{
				Text:  "",
				Style: "",
			})
			return emptyRet, fmt.Errorf("error running the OCR: %v", err)
		}
		// append the ocr output to styled text.
		styledOutput = append(styledOutput, common.StyledText{
			Text:  output,
			Style: "",
		})
	}

	// return the final styled output.
	return styledOutput, nil
}

// function to perform ocr using tesseract. returns raw text and an error.
func doOCR(imagePath string) (string, error) {
	// print message indicating ocr process.
	fmt.Println("Performing OCR on image...")

	// create a command to run tesseract with the image path.
	cmd := exec.Command("tesseract", imagePath, "stdout")
	// execute the command and capture output.
	output, err := cmd.Output()
	// handle any errors in the execution.
	if err != nil {
		return "", fmt.Errorf("error running Tesseract: %v", err)
	}

	// return the trimmed output text from tesseract.
	return strings.TrimSpace(string(output)), nil
}

// function to perform hocr using tesseract. returns styled text and an error.
func dohOCR(imagePath string) ([]common.StyledText, error) {
	// print message indicating hocr process.
	fmt.Println("Performing hOCR on image...")

	// create a command to run tesseract with hocr option.
	cmd := exec.Command("tesseract", imagePath, "stdout", "hocr")

	// execute the command and capture output.
	output, err := cmd.Output()
	// handle any errors in the execution.
	if err != nil {
		// return an empty styled text with error.
		var emptyRet []common.StyledText
		emptyRet = append(emptyRet, common.StyledText{
			Text:  "",
			Style: "",
		})

		return emptyRet, fmt.Errorf("error running Tesseract: %v", err)
	}

	// parse the hocr output and handle any errors.
	parsedOutput, err := parseHOCR(string(output))
	if err != nil {
		// return an empty styled text with error.
		var emptyRet []common.StyledText
		emptyRet = append(emptyRet, common.StyledText{
			Text:  "",
			Style: "",
		})
		return emptyRet, fmt.Errorf("%v", err)
	}

	// return the parsed hocr output as styled text.
	return parsedOutput, nil
}
