package main

import (
	"fmt"
	"log"
	"os"
	//path to dirs
	/*
		"photo-to-md/dependancies"
		"photo-to-md/save"
		"photo-to-md/capture"
		"photo-to-md/ocr"
	*/)

func main() {
	isInDocker := os.Getenv("INDOCKER")
	if isInDocker != "true" {
		dependancies.checkAndInstallDependencies()
	}

	imagePath := capture.captureImage()

	text, err := ocr.performOCR(imagePath)
	if err != nil {
		log.Fatalf("Failed to perform OCR: %v", err)
	}

	err = save.saveMarkdown(text)
	if err != nil {
		log.Fatalf("Failed to save markdown: %v", err)
	}

	fmt.Println("Markdown file created successfully!")
}
