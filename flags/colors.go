package flags

// colors of cmd

type ColorCMD = map[string]string

func getColorMap() ColorCMD {

	colorMap := ColorCMD{
		"black":       "0",
		"blue":        "1",
		"green":       "2",
		"aqua":        "3",
		"red":         "4",
		"purple":      "5",
		"yellow":      "6",
		"white":       "7",
		"gray":        "8",
		"lightblue":   "9",
		"lightgreen":  "A",
		"lightaqua":   "B",
		"lightred":    "C",
		"lightpurple": "D",
		"lightyellow": "E",
		"brightwhite": "F",
	}

	return colorMap

}
