package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	//path to dirs
	_ "photo-to-md/capture"
	"photo-to-md/dependencies"
	"photo-to-md/ocr"
	"photo-to-md/save"
)

func main() {

	mode := flag.String("mode", "path", "Mode: new capture or path (only path in this version)")
	path := flag.String("path", "", "Path to the image")
	specific := flag.Bool("s", false, "Fonts and styles included")
	flag.Parse()

	if *mode == "path" && *path == "" {
		fmt.Println("What are you doing with this flags...")
		flag.Usage()
		os.Exit(1)
	}

	if *specific == true {
		err := os.Setenv("ISSPECIFIC", "true")
		fmt.Println("Doing Specific check")
		if err != nil {
			log.Fatalf("Failed to set the ISSPECIFIC env flag: %v", err)
		}
	}

	var imagePath string

	if *mode == "capture" {
		/*
			imagePath = capture.CaptureImage()
		*/
		log.Fatalf("Capture option is not supported in this version")
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
