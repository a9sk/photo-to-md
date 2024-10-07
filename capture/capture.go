package capture

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func CaptureImage() string {
	imagePath := filepath.Join(os.TempDir(), "capture.jpg")
	cmd := exec.Command("ffmpeg", "-f", "video4linux2", "-i", "/dev/video0", "-frames:v", "1", imagePath)

	fmt.Println("Capturing image...")

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to capture image: %v", err)
	}

	fmt.Printf("Image captured and saved to %s\n", imagePath)
	return imagePath
}
