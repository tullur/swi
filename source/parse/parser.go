package parse

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"swi/source/model"
	"swi/source/util"

	"github.com/dlclark/regexp2"
)

func readXMLFile(fileName string) *os.File {
	file, err := os.Open(fileName + ".xml")
	util.CheckError(err)

	return file
}

// check element validity and rewrite input.xml
func checkValidity(name string) {
	file := readXMLFile(name)
	input, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	reObject := regexp2.MustCompile(`<\/*(object|field)>`, 0)
	reField := regexp2.MustCompile(`(?<=<([A-Za-z].*?)>).+?(?=</(\1.*?)>)`, 0)

	for i, line := range lines {
		str := (strings.Replace(line, " ", "", -1))
		if isMatch, _ := reObject.MatchString(str); isMatch {
			lines[i] = line
		} else if isMatch, _ := reField.MatchString(str); !isMatch {
			lines[i] = ""
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("valid.xml", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

// XMLtoJSON decode xml file and returns decoded data
func XMLtoJSON(name string) map[string]map[string]interface{} {
	checkValidity(name)
	defer os.Remove("valid.xml")

	file := readXMLFile("valid")
	defer file.Close()

	decoder := xml.NewDecoder(file)
	resultData := map[string]map[string]interface{}{}
	decoder.Strict = false

	for {
		object := model.XMLObject{}
		// read tokens from xml
		token, err := decoder.Token()

		if err != nil {
			if err == io.EOF {
				log.Println("Parsed")
				break
			} else {
				log.Println(err)
				break
			}
		}
		//  inspect token type
		switch se := token.(type) {
		case xml.StartElement:
			// if StartElement token has name object
			if se.Name.Local == "object" {
				// decode whole object element
				decoder.DecodeElement(&object, &se)

				// if <obj_name> ain't empty
				if object.ObjectName != "" {
					// write data to map
					resultData[object.ObjectName] = map[string]interface{}{}

					for _, val := range object.Fields {
						// check value type
						switch val.Type {
						case "int":
							intValue, err := strconv.ParseInt(val.Value, 0, 32)
							util.CheckError(err)
							if intValue != 0 {
								resultData[object.ObjectName][val.Name] = intValue
							} else {
								delete(resultData[object.ObjectName], val.Value)
							}
						case "string":
							if val.Value != "" {
								resultData[object.ObjectName][val.Name] = val.Value
							} else {
								delete(resultData[object.ObjectName], val.Value)
							}
						default:
							delete(resultData[object.ObjectName], val.Value)
						}
					}
				}
			}
		}
	}

	return resultData
}
