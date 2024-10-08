package save

import (
	"fmt"
	"os"
	"time"
)

type StyledText struct {
	Text  string
	Style string
}

func SaveMarkdown(text []StyledText) error {

	var stringText string

	if os.Getenv("ISSPECIFIC") == "true" {
		stringText = htmlToMarkdown(text)
	} else {
		stringText = text[0].Text
	}

	markdownContent := fmt.Sprintf("# Extracted Text - %s\n\n%s", time.Now().Format(time.RFC1123), stringText)

	outputPath := "output.md"
	err := os.WriteFile(outputPath, []byte(markdownContent), 0644)
	if err != nil {
		return fmt.Errorf("error saving markdown file: %v", err)
	}

	fmt.Printf("Markdown saved to: %s\n", outputPath)
	return nil
}
