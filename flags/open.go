package flags

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

	var batchStr strings.Builder
	batchStr.WriteString("@echo off \n")

	color := "00"

	if args.Colour != "" {
		color = fmt.Sprintf("color %s%s \n", getColorMap()[args.Colour], "1")
		batchStr.WriteString(color)
	}

	strData := string(data)
	strData = batchStr.String() + strData

	fmt.Println(strData)

	err = os.WriteFile(batchFileName, []byte(strData), 0777)

	if err != nil {
		fmt.Println("Error writing the file")
	}

	// Start a new CMD process to run the batch file
	cmd := exec.Command("cmd", "/C", "start", batchFileName)
	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

}
