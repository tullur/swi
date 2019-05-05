package parse

import (
	"encoding/xml"
	"os"
	"strconv"
	"swi/source/model"
	"swi/source/util"
)

// XMLtoJSON decode xml file and returns decoded data
func XMLtoJSON(file *os.File) map[string]map[string]interface{} {
	decoder := *xml.NewDecoder(file)
	resultData := map[string]map[string]interface{}{}

	for {
		// read tokens from xml
		t, err := decoder.Token()
		util.CheckError(err)
		if t == nil {
			break
		}
		object := model.XMLObject{}

		//  inspect token type
		switch se := t.(type) {
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
							resultData[object.ObjectName][val.Name] = "Incorrect Data Type"
						}
					}
				}
			}
		}
	}

	return resultData
}
