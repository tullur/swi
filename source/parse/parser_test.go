package parse

import (
	"reflect"
	"testing"
)

func TestXMLtoJSON(t *testing.T) {
	wantData := XMLtoJSON("../data/input")
	tag := XMLtoJSON("../data/tag")

	tests := []struct {
		name string
		arg  string
		want map[string]map[string]interface{}
	}{
		{"XML with root", "../data/test/test_root", wantData},
		{"XML without root", "../data/test/test_without_root", wantData},
		{"XML invalid tag", "../data/test/test_with_tag", tag},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XMLtoJSON(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XMLtoJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
