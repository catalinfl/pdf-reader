package process

import (
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

		text, _ := extractTextFromPage(&page)

		// if err != nil {
		// 	return err
		// }

		if i != r.NumPage() {
			finalText += text + "\n"
		} else {
			finalText += text
		}

	}

	return finalText, nil

}

func extractTextFromPage(page *pdf.Page) (string, error) {
	var text string = ""

	content := page.Content()

	// verify := page.Content().Text
	// fmt.Println(verify)
	tolerance := 1.0

	for i := 1; i < len(content.Text); i++ {

		startOfThis := content.Text[i].X
		endOfLast := content.Text[i-1].X + content.Text[i-1].W

		endOfLastY := content.Text[i-1].Y
		currentY := content.Text[i].Y

		if endOfLast <= startOfThis-tolerance || endOfLast >= startOfThis+tolerance {
			text += content.Text[i-1].S
			text += " "
		} else {
			text += content.Text[i-1].S
		}

		if currentY <= endOfLastY-tolerance || currentY >= endOfLastY+tolerance {
			text += "\n"
		}

		// if t == y {
		// 	fmt.Println("Sunt egale")
		// }

		// fmt.Println(t.FontSize)
	}

	text += content.Text[len(content.Text)-1].S
	// fmt.Printf("%s", text)

	return text, nil
}
