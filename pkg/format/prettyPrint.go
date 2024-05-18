package format

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/xemkayx/steam-api/pkg/steamclient/config"
)

// todo: pretty print for vdf

// PrettyPrints a returned object in the given format
func PrettyPrint(v interface{}, format config.OutputFormat) string {
	switch format {
	case config.Json:
		return prettyPrintJSON(v)
	case config.Xml:
		return prettyPrintXML(v)
	default:
		return fmt.Sprintf("Unsupported format: %s", format)
	}
}

// pretty-prints a JSON object
func prettyPrintJSON(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error printing object as JSON: %v", err)
	}
	return string(data)
}

// pretty-prints an XML object
func prettyPrintXML(v interface{}) string {
	data, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error printing object as XML: %v", err)
	}
	return string(data)
}
