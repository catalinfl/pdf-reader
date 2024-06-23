package flags

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sync"
)

// de adaugat background color

func OpenCMD(args Arguments, wg *sync.WaitGroup) {
	defer wg.Done()

	batchFileName := "tempInput.bat"

	data, err := os.ReadFile(batchFileName)

	if err != nil {
		fmt.Println("Error reading the file")
	}

	strData := string(data)

	colorsMap := getColorMap() // get colors from map
	newColorCmd := fmt.Sprintf("color %s%s", colorsMap[args.Background], colorsMap[args.Colour])

	colorCmdRegex := regexp.MustCompile(`color \w\w`)

	if colorCmdRegex.MatchString(strData) {
		strData = colorCmdRegex.ReplaceAllString(strData, newColorCmd)
	} else {
		strData = "@echo off \n" + newColorCmd + strData
	}

	err = os.WriteFile(batchFileName, []byte(strData), 0777)

	if err != nil {
		fmt.Println("Error writing the file")
		return
	}

	// Start a new CMD process to run the batch file
	cmd := exec.Command("cmd", "/C", "start", batchFileName)
	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

}
