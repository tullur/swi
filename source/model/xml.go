package model

import "encoding/xml"

// XMLObject ->
type XMLObject struct {
	XMLName    xml.Name   `xml:"object"`
	ObjectName string     `xml:"obj_name,omitempty"`
	Fields     []XMLField `xml:"field,omitempty"`
}

// XMLField ->
type XMLField struct {
	Name  string `xml:"name,omitempty"`
	Type  string `xml:"type,omitempty"`
	Value string `xml:"value,omitempty"`
}
