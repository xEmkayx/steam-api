package format

import (
	"fmt"
	"testing"

	"github.com/xemkayx/steam-api/pkg/steamclient/config"
)

type SimpleStruct struct {
	Name string
	Age  int
}

type NestedStruct struct {
	Name  string
	Age   int
	Inner SimpleStruct
}

type Platform struct {
	Name string `xml:"name"`
	Type string `xml:"type"`
}

func TestPrettyPrintJSON(t *testing.T) {
	input := map[string]interface{}{
		"name": "Steam",
		"type": "Platform",
	}
	expectedOutput := `{
  "name": "Steam",
  "type": "Platform"
}`
	output := PrettyPrint(input, config.Json)
	if output != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output)
	}
}

func TestPrettyPrintXML(t *testing.T) {
	input := Platform{
		Name: "Steam",
		Type: "Platform",
	}

	expectedOutput := `<Platform>
  <name>Steam</name>
  <type>Platform</type>
</Platform>`
	output := PrettyPrint(input, config.Xml)
	if output != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output)
	}
}

func TestPrettyPrintUnsupportedFormat(t *testing.T) {
	input := map[string]interface{}{
		"name": "Steam",
		"type": "Platform",
	}

	expectedErrorMessage := fmt.Sprintf("Unsupported format: %s", config.OutputFormat(999))
	output := PrettyPrint(input, config.OutputFormat(999))
	if output != expectedErrorMessage {
		t.Errorf("Expected error message: %s, got: %s", expectedErrorMessage, output)
	}
}

func TestPrettyPrintErrors(t *testing.T) {
	tests := []struct {
		name   string
		format config.OutputFormat
		want   string
	}{
		{
			name:   "unknown format",
			format: config.OutputFormat(4),
			want:   "Unsupported format: unknown format",
		},
		// Add more test cases here for other error scenarios
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrettyPrint(nil, tt.format); got != tt.want {
				t.Errorf("PrettyPrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
