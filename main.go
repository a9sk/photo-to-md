package main

import (
	"fmt"
	"log"
	"os"

	//path to dirs
	"photo-to-md/capture"
	"photo-to-md/dependencies"
	"photo-to-md/ocr"
	"photo-to-md/save"
)

func main() {
	isInDocker := os.Getenv("INDOCKER")
	if isInDocker != "true" {
		dependencies.CheckAndInstallDependencies()
	}

	imagePath := capture.CaptureImage()

	text, err := ocr.PerformOCR(imagePath)
	if err != nil {
		log.Fatalf("Failed to perform OCR: %v", err)
	}

	err = save.SaveMarkdown(text)
	if err != nil {
		log.Fatalf("Failed to save markdown: %v", err)
	}

	fmt.Println("Markdown file created successfully!")
}
