package save

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func SaveMarkdown(text string) error {
	markdownContent := fmt.Sprintf("# Extracted Text - %s\n\n%s", time.Now().Format(time.RFC1123), text)

	outputPath := filepath.Join(os.TempDir(), "output.md")
	err := os.WriteFile(outputPath, []byte(markdownContent), 0644)
	if err != nil {
		return fmt.Errorf("error saving markdown file: %v", err)
	}

	fmt.Printf("Markdown saved to: %s\n", outputPath)
	return nil
}
