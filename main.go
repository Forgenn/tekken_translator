package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		// Set logger
		l := log.New(log.Writer(), log.Prefix(), log.Flags())

		// Create astilectron
		a, err := astilectron.New(l, astilectron.Options{
			AppName:           "Test",
			BaseDirectoryPath: "resources",
		})
		if err != nil {
			l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
		}
		defer a.Close()

		// Handle signals
		a.HandleSignals()

		// Start
		if err = a.Start(); err != nil {
			l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
		}

		// New window
		var w *astilectron.Window
		if w, err = a.NewWindow("resources/index.html", &astilectron.WindowOptions{
			Center: astikit.BoolPtr(true),
			Height: astikit.IntPtr(700),
			Width:  astikit.IntPtr(700),
		}); err != nil {
			l.Fatal(fmt.Errorf("main: new window failed: %w", err))
		}
	*/
	initDict()
	var move = "1+2"
	translate(move)
	/*
		// Create windows
		if err = w.Create(); err != nil {
			l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
		}

		// Blocking pattern
		a.Wait()
	*/
}

var stringToImageDict = map[string]string{}

func initDict() {
	stringToImageDict["1"] = "1.png"
	stringToImageDict["1+2"] = "1+2.png"
	stringToImageDict["1+3"] = "1+3.png"
	stringToImageDict["1+4"] = "1+4.png"
	stringToImageDict["2"] = "2.png"
	stringToImageDict["2+4"] = "2+4.png"
	stringToImageDict["3"] = "3.png"
	stringToImageDict["3+4"] = "3+4.png"
	stringToImageDict["b"] = "b.png"
	stringToImageDict["bp"] = "bp.png"
	stringToImageDict["d"] = "d.png"
	stringToImageDict["df"] = "df.png"
	stringToImageDict["dp"] = "dp.png"
	stringToImageDict["f"] = "f.png"
	stringToImageDict["fp"] = "fp.png"
	stringToImageDict["n"] = "n.png"
	stringToImageDict["u"] = "u.png"
	stringToImageDict["uf"] = "uf.png"
	stringToImageDict["ub"] = "ub.png"
}

func translate(moveString string) {

	imageMoveString := []string{}
	var backPointer = 0

	for i := 0; i < len(moveString); i++ {
		if string(moveString[i]) == "+" {
			if _, err := strconv.Atoi(string(moveString[i-1])); err == nil {
				if _, err2 := strconv.Atoi(string(moveString[i+1])); err2 == nil {
					imageMoveString = append(imageMoveString, stringToImageDict[moveString[i-1:i+2]])
					backPointer = i + 2
				}
			} else {
				imageMoveString = append(imageMoveString, stringToImageDict[moveString[backPointer:i]])
				backPointer = i + 1
			}
		}

		if string(moveString[i]) == "," {
			imageMoveString = append(imageMoveString, stringToImageDict[moveString[backPointer:i]])
			backPointer = i + 1
		}

		if i == len(moveString)-1 {
			if val, ok := stringToImageDict[moveString[backPointer:]]; ok {
				imageMoveString = append(imageMoveString, val)
			}
		}
	}
	fmt.Println(imageMoveString)
}
