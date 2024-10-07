package ocr

import (
	"fmt"
	"os/exec"
	"strings"
)

func PerformOCR(imagePath string) (string, error) {
	fmt.Println("Performing OCR on image...")

	cmd := exec.Command("tesseract", imagePath, "stdout")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running Tesseract: %v", err)
	}

	return strings.TrimSpace(string(output)), nil
}
