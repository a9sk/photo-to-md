package dependencies

import (
	"fmt"
	"log"
	"os/exec"
)

func installTesseract() {
	fmt.Println("Installing Tesseract...")
	var cmd *exec.Cmd

	switch detectOS() {
	case "linux":
		cmd = exec.Command("sudo", "apt-get", "install", "-y", "tesseract-ocr")
	case "darwin":
		cmd = exec.Command("brew", "install", "tesseract")
	case "windows":
		fmt.Println("Please install Tesseract manually from https://github.com/tesseract-ocr/tesseract/wiki")
		log.Fatalf("Install Tesseract and run the code.")
		return
	default:
		fmt.Println("Unsupported OS. Please install Tesseract manually.")
		return
	}

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to install Tesseract: %v", err)
	}
	fmt.Println("Tesseract installed successfully!")
}
