package save

import (
	"fmt"
	"os"
	_ "time"

	"github.com/a9sk/photo-to-md/common"
)

func SaveMarkdown(text []common.StyledText) error {

	var stringText string

	if os.Getenv("ISSPECIFIC") == "true" {
		stringText = htmlToMarkdown(text)
	} else {
		stringText = text[0].Text
	}

	// to echo in the file some info about extraction uncomment following line and comment the one after
	// markdownContent := fmt.Sprintf("# Extracted Text - %s\n\n%s", time.Now().Format(time.RFC1123), stringText)

	markdownContent := stringText

	outputPath := "output.md"
	err := os.WriteFile(outputPath, []byte(markdownContent), 0644)
	if err != nil {
		return fmt.Errorf("error saving markdown file: %v", err)
	}

	fmt.Printf("Markdown saved to: %s\n", outputPath)
	return nil
}
