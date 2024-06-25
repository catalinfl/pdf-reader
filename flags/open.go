package flags

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/catalinfl/pdfreader/process"
)

func OpenCMD(args Arguments, wg *sync.WaitGroup) {
	defer wg.Done()

	batchFileName := "tempInput.bat"

	// Create a new batch file

	batchText := `@echo off 
	cls
	color 07`

	_, err := os.Stat(batchFileName)

	// Read the existing content of the batch file
	if os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating it...")
		// File does not exist, create it with the batchText content
		err = os.WriteFile(batchFileName, []byte(batchText), 0666)
		if err != nil {
			fmt.Println("Error writing the file:", err)
			return
		}
		fmt.Println("File created successfully.")
	} else if err != nil {
		fmt.Println("Error checking the file:", err)
		return
	} else {
		fmt.Println("File already exists.")
	}

	strData := batchText

	// extract pdf
	text, err := process.ExtractTextFromPDF(args.ReadPath)
	if err != nil {
		fmt.Println("Error extracting text from PDF:", err)
		return
	}

	// verify if color exists in batch, if exists it modifies, if not it adds it
	colorsMap := getColorMap()
	newColorCmd := fmt.Sprintf("color %s%s", colorsMap[args.Background], colorsMap[args.Colour])
	colorCmdRegex := regexp.MustCompile(`color \w\w`)
	strData = colorCmdRegex.ReplaceAllString(strData, newColorCmd)

	// displaying lines logic
	lines := strings.Split(text, "\n")
	totalLines := len(lines)
	linesPerPage := 27
	totalPages := (totalLines + linesPerPage - 1) / linesPerPage // Ceiling division

	// write text in a batch file
	err = os.WriteFile(batchFileName, []byte(strData), 0666)
	if err != nil {
		fmt.Println("Error writing the file:", err)
		return
	}

	// start a new cmd process
	cmd := exec.Command("cmd", "/C", batchFileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing the batch file:", err)
		return
	}

	// add reader
	reader := bufio.NewReader(os.Stdin)
	currentPage := 1

	for {
		// for starts for every page
		clearScreen()

		startLine := (currentPage - 1) * linesPerPage
		endLine := startLine + linesPerPage
		if endLine > totalLines {
			endLine = totalLines
		}

		for _, line := range lines[startLine:endLine] {
			fmt.Println(line)
		}

		for i := 0; i < linesPerPage-(endLine-startLine); i++ {
			fmt.Println()
		}

		fmt.Printf("\n<< < Page %d/%d > >> q - quit, goto [page]\n", currentPage, totalPages)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == ">" && currentPage < totalPages {
			currentPage++
		} else if input == "<" && currentPage > 1 {
			currentPage--
		} else if input == "q" {
			os.Remove(batchFileName)
			break
		} else if strings.Contains(input, "goto") {
			// get page number
			pageNumber := strings.Split(input, " ")[1]
			intPageNumber, _ := strconv.Atoi(pageNumber)
			if intPageNumber < 0 {
				currentPage = 0
			}

			if intPageNumber > totalPages {
				currentPage = totalPages
			}

			if intPageNumber > 0 && intPageNumber <= totalPages {
				currentPage = intPageNumber
			}
		}
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
