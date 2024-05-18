package steamclient

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"steam-api/pkg/steamclient/config"
)

const apiKeyErrorMessage = "you have to specify an API-key to call this endpoint"

func decodeResponse(format config.OutputFormat, body io.ReadCloser, result interface{}) (interface{}, error) {
	switch format {
	case config.Json:
		if _, err := decodeJSON(&result, body); err != nil {
			return nil, err
		}
		return result, nil

	case config.Xml:
		if _, err := decodeXML(&result, body); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, fmt.Errorf("unsupported format requested: %v", format)
	}
}

func decode[T any](dest *T, r io.Reader, decoderFunc func(interface{}) error) (*T, error) {
	if err := decoderFunc(dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// decodeJSON dekodiert JSON-Daten aus einem Reader in das Zielobjekt dest.
func decodeJSON[T any](dest *T, r io.Reader) (*T, error) {
	jsonDecoder := json.NewDecoder(r)
	return decode(dest, r, jsonDecoder.Decode)
}

// decodeXML dekodiert XML-Daten aus einem Reader in das Zielobjekt dest.
func decodeXML[T any](dest *T, r io.Reader) (*T, error) {
	xmlDecoder := xml.NewDecoder(r)
	return decode(dest, r, xmlDecoder.Decode)
}
