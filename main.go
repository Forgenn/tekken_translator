package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

var stringToImageDict = map[string]string{}

func main() {

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

	initDict()
	var move = "uf+3,4"
	translate(move)

	// Create windows
	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}
	w.OpenDevTools()
	w.SendMessage("hello", func(m *astilectron.EventMessage) {
		// Unmarshal
		var s string
		m.Unmarshal(&s)

		// Process message
		log.Printf("received %s\n", s)
	})
	// Blocking pattern
	a.Wait()

}

func initDict() {
	stringToImageDict["1"] = "1.svg"
	stringToImageDict["1+2"] = "1+2.svg"
	stringToImageDict["1+3"] = "1+3.svg"
	stringToImageDict["1+4"] = "1+4.svg"
	stringToImageDict["2"] = "2.svg"
	stringToImageDict["2+4"] = "2+4.svg"
	stringToImageDict["3"] = "3.svg"
	stringToImageDict["3+4"] = "3+4.svg"
	stringToImageDict["4"] = "4.svg"
	stringToImageDict["b"] = "b.svg"
	stringToImageDict["bp"] = "bp.svg"
	stringToImageDict["d"] = "d.svg"
	stringToImageDict["df"] = "df.svg"
	stringToImageDict["dp"] = "dp.svg"
	stringToImageDict["f"] = "f.svg"
	stringToImageDict["fp"] = "fp.svg"
	stringToImageDict["n"] = "n.svg"
	stringToImageDict["u"] = "u.svg"
	stringToImageDict["uf"] = "uf.svg"
	stringToImageDict["ub"] = "ub.svg"
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
