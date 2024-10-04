package dependencies

import (
	"fmt"
	"log"
	"os/exec"
)

func installFFmpeg() {
	fmt.Println("Installing ffmpeg...")
	var cmd *exec.Cmd

	switch detectOS() {
	case "linux":
		cmd = exec.Command("sudo", "apt-get", "install", "-y", "ffmpeg")
	case "darwin":
		cmd = exec.Command("brew", "install", "ffmpeg")
	case "windows":
		fmt.Println("Please install ffmpeg manually from https://ffmpeg.org/download.html")
		log.Fatalf("Install ffmpeg and run the code.")
	default:
		fmt.Println("Unsupported OS. Please install ffmpeg manually.")
		return
	}

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to install ffmpeg: %v", err)
	}
	fmt.Println("ffmpeg installed successfully!")
}
