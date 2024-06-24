package flags

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"

	"github.com/catalinfl/pdfreader/process"
)

func OpenCMD(args Arguments, wg *sync.WaitGroup) {
	defer wg.Done()

	batchFileName := "tempInput.bat"

	// Read the existing content of the batch file
	data, err := os.ReadFile(batchFileName)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	strData := string(data)

	// Extract text from PDF
	text, err := process.ExtractTextFromPDF(args.ReadPath)
	if err != nil {
		fmt.Println("Error extracting text from PDF:", err)
		return
	}

	lines := strings.Split(text, "\n")
	totalLines := len(lines)
	linesPerPage := 25
	totalPages := (totalLines + linesPerPage - 1) / linesPerPage // Ceiling division

	// Modify the batch file content here if needed, but remove pagination logic

	colorsMap := getColorMap()
	newColorCmd := fmt.Sprintf("color %s%s", colorsMap[args.Background], colorsMap[args.Colour])

	colorCmdRegex := regexp.MustCompile(`color \w\w`)
	strData = colorCmdRegex.ReplaceAllString(strData, newColorCmd)

	// Write the modified content back to the batch file with more restrictive permissions
	err = os.WriteFile(batchFileName, []byte(strData), 0666)
	if err != nil {
		fmt.Println("Error writing the file:", err)
		return
	}

	// Start a new CMD process to run the batch file
	cmd := exec.Command("cmd", "/C", batchFileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing the batch file:", err)
		return
	}

	// Handle pagination in Go
	reader := bufio.NewReader(os.Stdin)
	currentPage := 1

	for {

		clearScreen()

		startLine := (currentPage - 1) * linesPerPage
		endLine := startLine + linesPerPage
		if endLine > totalLines {
			endLine = totalLines
		}
		for _, line := range lines[startLine:endLine] {
			fmt.Println(line)
		}

		// for i := 0; i < 3; i++ {
		// 	fmt.Println()
		// }

		for i := 0; i < linesPerPage-(endLine-startLine)+2; i++ {
			fmt.Println()
		}

		fmt.Printf("\n<< < Page %d/%d > >> q to quit\n", currentPage, totalPages)
		// fmt.Println("Press 'n' for next page, 'p' for previous page, or 'q' to quit.")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == ">" && currentPage < totalPages {
			currentPage++
		} else if input == "<" && currentPage > 1 {
			currentPage--
		} else if input == "q" {
			break
		}
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
