package main

import (
	"log"
	"os"
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
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("SWI STARTED")
}

func main() {
	file, err := os.Open(inputFileName + ".xml")
	util.CheckError(err)

	defer file.Close()
	resultData := parse.XMLtoJSON(file)

	util.WriteJSON(outputFileName, resultData)
}
