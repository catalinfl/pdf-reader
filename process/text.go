package process

import (
	"strings"

	"rsc.io/pdf"
)

func ExtractTextFromPDF(filePath string) (string, error) {

	r, err := pdf.Open(filePath)

	if err != nil {
		return "", err
	}

	finalText := ""

	for i := 1; i <= r.NumPage(); i++ {
		page := r.Page(i)

		if page.V.IsNull() {
			continue
		}

		text, err := extractTextFromPage(&page)

		if err != nil {
			return "", err
		}

		finalText += text

		if i != r.NumPage() {
			finalText += "\n"
		}
	}

	return finalText, nil

}

func extractTextFromPage(page *pdf.Page) (string, error) {
	var text string = ""

	content := page.Content()

	// add a tolerance for x coordinate
	tolerance := 1.0

	for i := 1; i < len(content.Text); i++ {

		startOfThis := content.Text[i].X
		endOfLast := content.Text[i-1].X + content.Text[i-1].W

		endOfLastY := content.Text[i-1].Y
		currentY := content.Text[i].Y

		text += removeNonASCII(content.Text[i-1].S)

		// if the x coordinate of the current letter + width is not close to the x coordinate of the last letter put a space between
		if endOfLast <= startOfThis-tolerance || endOfLast >= startOfThis+tolerance {
			text += " "
		}

		// same for y, but new line
		if currentY <= endOfLastY-tolerance || currentY >= endOfLastY+tolerance {
			text += "\n"
		}
	}

	text += content.Text[len(content.Text)-1].S

	lines := strings.Split(text, "\n")

	// remove useless lines
	for i := len(lines) - 1; i >= 0; i-- {
		lines[i] = strings.TrimSpace(lines[i])

		if len(lines[i]) < 3 {

			// delete the current line if there are less than 3 characters
			lines = append(lines[:i], lines[i+1:]...)
		}
	}

	text = strings.Join(lines, "\n")

	return text, nil
}

func removeNonASCII(s string) string {
	result := ""
	for _, c := range s {
		if c < 128 {
			result += string(c)
		}
	}
	return result
}
