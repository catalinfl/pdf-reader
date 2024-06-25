package flags

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
)

type Arguments struct {
	Colour     string
	ReadPath   string
	Background string
	Help       bool
}

var args Arguments

func ArgumentsFunc() {

	flag.StringVar(&args.Colour, "colour", "white", "colour of cli text")
	flag.StringVar(&args.ReadPath, "read", "", "read the path")
	flag.StringVar(&args.Background, "background", "black", "background of cli")
	flag.BoolVar(&args.Help, "help", false, "help")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Use --help to see all commands available.")
	}

	flag.Parse()

	stringArg, errorString := CheckArguments(args)

	if errorString != "" {
		fmt.Printf("%s", errorString)
		return
	}

	fmt.Printf("%s", stringArg)

	if args.Help {
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go OpenCMD(args, &wg)

	wg.Wait()

}

func CheckArguments(args Arguments) (string, string) {

	val := reflect.ValueOf(args)
	var info, err strings.Builder

	for i := 0; i < val.NumField(); i++ {

		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		switch field.Name {
		case "ReadPath":
			if value == "" {
				err.WriteString("You need to set the read path! \nUse --read [path] to set the path \n")
				args = Arguments{Colour: "", ReadPath: "", Background: "", Help: false}
			} else {
				info.WriteString(fmt.Sprintf("Read path is %s. \n", args.ReadPath))
			}
		case "Help":
			if value == true {
				info.Reset()
				info.WriteString("\n--colour [colour] - Set the colour of CLI\n--read [path] - Read the path of .pdf \n--background [colour] - Set background colour \n\nType > - go to next page\nType < - go to previous page\nType q - quit\nType goto [page number] - go to specific page number\n")
				err.Reset()
				return info.String(), err.String()
			}
		case "Colour":
			if value == "" {
				err.WriteString("Colour is not set \n")
			} else {
				info.WriteString(fmt.Sprintf("Colour is set as %s. \n", args.Colour))
			}
		case "Background":
			info.WriteString(fmt.Sprintf("Background set as %s \n", args.Background))
		}
	}

	return info.String(), err.String()

}
