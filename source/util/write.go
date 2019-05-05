package util

import (
	"encoding/json"
	"os"
)

// WriteJSON write unnmarshaled xml data to the JSON file
func WriteJSON(name string, data map[string]map[string]interface{}) {
	f, err := os.Create(name + ".json")
	CheckError(err)

	result, err := json.MarshalIndent(data, "", "\t")
	CheckError(err)

	f.Write(result)
}
