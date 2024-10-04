package dependencies

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func CheckAndInstallDependencies() {
	fmt.Println("Checking for required dependencies...")

	if !isCommandAvailable("ffmpeg") {
		fmt.Println("ffmpeg is not installed.")
		installTool("ffmpeg")
	} else {
		fmt.Println("ffmpeg is installed.")
	}

	if !isCommandAvailable("tesseract") {
		fmt.Println("Tesseract is not installed.")
		installTool("tesseract")
	} else {
		fmt.Println("Tesseract is installed.")
	}
}

func isCommandAvailable(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func installTool(toolName string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Do you want to install %s? (yes/no): ", toolName)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	if response == "yes" || response == "y" {
		switch toolName {
		case "ffmpeg":
			installFFmpeg()
		case "tesseract":
			installTesseract()
		default:
			fmt.Printf("Unknown tool: %s\n", toolName)
		}
	} else {
		log.Fatalf("%s is required. Exiting...\n", toolName)
	}
}

func detectOS() string {
	switch os := runtime.GOOS; os {
	case "linux":
		return "linux"
	case "darwin":
		return "darwin"
	case "windows":
		return "windows"
	default:
		return "unknown"
	}
}
