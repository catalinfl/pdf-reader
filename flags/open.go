package flags

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"

	"github.com/catalinfl/pdfreader/process"
)

// de adaugat background color

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

	startMarker := ":: START ECHO COMMANDS"
	endMarker := ":: END ECHO COMMANDS"

	// Extract text from PDF
	text, err := process.ExtractTextFromPDF(args.ReadPath)

	if err != nil {
		fmt.Println("Error extracting text from PDF:", err)
		return
	}

	lines := strings.Split(text, "\n")

	sectionRegex := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(startMarker) + `.*` + regexp.QuoteMeta(endMarker))

	newEchoCmds := startMarker + "\n"

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			newEchoCmds += "echo. \n" // Echo a blank line if line is empty
		} else {
			newEchoCmds += fmt.Sprintf("echo %s \n", lines[i])
		}
	}

	newEchoCmds += endMarker

	strData = sectionRegex.ReplaceAllString(strData, newEchoCmds)

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
	cmd := exec.Command("cmd", "/C", "start", batchFileName)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing the batch file:", err)
	}
}
