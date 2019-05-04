package parser

import (
	"encoding/xml"
	"os"
	"strconv"
	"swi/source/model"
	"swi/source/utils"
)

func XMLtoJSON(name string) {
	file, err := os.Open(name + ".xml")
	utils.CheckError(err)

	defer file.Close()

	decoder := *xml.NewDecoder(file)

	resultData := map[string]map[string]interface{}{}
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		object := model.XMLObject{}

		switch et := t.(type) {
		case xml.StartElement:
			if et.Name.Local == "object" {

				decoder.DecodeElement(&object, &et)
				if object.ObjectName != "" {
					resultData[object.ObjectName] = map[string]interface{}{}

					for _, val := range object.Fields {
						switch val.Type {
						case "int":
							intValue, _ := strconv.ParseInt(val.Value, 0, 32)
							resultData[object.ObjectName][val.Name] = intValue
						case "string":
							resultData[object.ObjectName][val.Name] = val.Value
						default:
							resultData[object.ObjectName][val.Name] = "Incorrect data type"
						}

					}
				}
			}
		}

		utils.WriteJSON("output", resultData)
	}
}
