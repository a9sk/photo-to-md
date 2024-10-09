package save

import (
	"photo-to-md/common"
	"strings"
)

func htmlToMarkdown(styledTexts []common.StyledText) string {
	var sb strings.Builder

	for _, styledText := range styledTexts {
		switch styledText.Style {
		case "bold":
			sb.WriteString("**" + styledText.Text + "**")
		case "italic":
			sb.WriteString("_" + styledText.Text + "_")
		case "title":
			sb.WriteString("# " + styledText.Text + "\n")
		case "semi-title":
			sb.WriteString("## " + styledText.Text + "\n")
		default:
			sb.WriteString(styledText.Text)
		}
	}

	return sb.String()
}
