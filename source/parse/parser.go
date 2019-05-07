package parse

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"strconv"
	"swi/source/model"
	"swi/source/util"
)

func readXMLFile(fileName string) *os.File {
	file, err := os.Open(fileName + ".xml")
	util.CheckError(err)

	return file
}

// XMLtoJSON decode xml file and returns decoded data
func XMLtoJSON(name string) map[string]map[string]interface{} {
	file := readXMLFile(name)
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
							resultData[object.ObjectName][val.Name] = intValue
						case "string":
							resultData[object.ObjectName][val.Name] = val.Value
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
