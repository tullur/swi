package model

import "encoding/xml"

// XMLObject desribe xml Object node
type XMLObject struct {
	XMLName    xml.Name   `xml:"object"`
	ObjectName string     `xml:"obj_name"`
	Fields     []XMLField `xml:"field"`
}

// XMLField describe xml Field node
type XMLField struct {
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Value string `xml:"value"`
}
