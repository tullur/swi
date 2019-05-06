package util

import (
	"bytes"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

func _UnescapeUnicodeCharactersInJSON(_jsonRaw json.RawMessage) (json.RawMessage, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(_jsonRaw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

// WriteJSON write unnmarshaled xml data to the JSON file
func WriteJSON(name string, data map[string]map[string]interface{}) {
	f, err := os.Create(name + ".json")
	CheckError(err)

	result, err := json.MarshalIndent(data, "", "\t")
	result = bytes.Replace(result, []byte("\\u003c"), []byte("<"), -1)
	result = bytes.Replace(result, []byte("\\u003e"), []byte(">"), -1)
	result = bytes.Replace(result, []byte("\\u0026"), []byte("&"), -1)

	CheckError(err)
	f.Write(result)
}
