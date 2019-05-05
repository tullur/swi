package main

import (
	"log"
	"swi/source/parse"
	"swi/source/util"
)

const (
	inputFileName  string = "input"
	outputFileName string = "output"
)

// write logs to console/terminal
func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("SWI STARTED")
}

func main() {
	resultData := parse.XMLtoJSON(inputFileName)

	util.WriteJSON(outputFileName, resultData)
}
