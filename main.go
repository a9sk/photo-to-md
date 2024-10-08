package main

import (
	"flag"
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

	mode := flag.String("mode", "capture", "Mode: new capture or path")
	path := flag.String("path", "", "Path to the image")
	flag.Parse()

	if *mode == "path" && *path == "" {
		fmt.Println("What are you doing with this flags...")
		flag.Usage()
		os.Exit(1)
	}

	var imagePath string

	if *mode == "capture" {
		imagePath = capture.CaptureImage()
	} else {
		imagePath = *path
	}

	isInDocker := os.Getenv("INDOCKER")
	if isInDocker != "true" {
		dependencies.CheckAndInstallDependencies()
	}

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
