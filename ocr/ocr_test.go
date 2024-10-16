package ocr

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/a9sk/photo-to-md/common"
)

var expected = []common.StyledText{
	{Text: "Title\n", Style: "title"},
	{Text: "Semi-title\n", Style: "semi-title"},
	{Text: "Single line of text...\n", Style: "normal"},
	{Text: "Multiple lines of text, this lines should be considered one. Second part\n", Style: "normal"},
	{Text: "of the multiple lines of text.\n", Style: "normal"},
	{Text: "Another title\n", Style: "title"},
	{Text: "mock_ocr_text.go\n", Style: "semi-title"},
	{Text: "\n", Style: "normal"},
}

// unit test for parseHOCR function
func TestParseHOCR(t *testing.T) {

	// convert the relative path to an absolute path
	imagePath, err := filepath.Abs("../testdata/images/test_hocr_image.png")
	if err != nil {
		t.Fatalf("error getting absolute path: %v", err)
	}

	actual, err := dohOCR(imagePath)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(expected) != len(actual) {
		t.Fatalf("lengths are not equal. Expected %d, got %d", len(expected), len(actual))
	}

	for i := range expected {
		expText := strings.TrimSpace(expected[i].Text)
		actText := strings.TrimSpace(actual[i].Text)

		if expText != actText || expected[i].Style != actual[i].Style {
			t.Errorf("mismatch at index %d:\nExpected: %+v\nGot: %+v", i, expected[i], actual[i])
		}
	}
}
