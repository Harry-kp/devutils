package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/gizak/termui/v3/widgets"
	"gopkg.in/yaml.v2"
)

// DecodeJWT takes a JWT string and returns the decoded header and payload.
func DecodeJWT(jwt string) (string, string, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return "", "", fmt.Errorf("Invalid JWT token format")
	}
	headerJSON, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return "", "", fmt.Errorf("Failed to decode header")
	}
	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", "", fmt.Errorf("Failed to decode payload")
	}
	return string(headerJSON), string(payloadJSON), nil
}

func ConvertYamlToJson(yamlStr string) (string, error) {
	// Convert YAML to a generic map
	var yamlData map[string]interface{}
	err := yaml.Unmarshal([]byte(yamlStr), &yamlData)
	if err != nil {
		return "", fmt.Errorf("Invalid YAML format")
	}

	// Convert the map to JSON
	jsonData, err := json.MarshalIndent(yamlData, "", "  ")
	if err != nil {
		return "", fmt.Errorf("Invalid YAML format")
	}

	pretty_json, err := FormatJSON(string(jsonData))
	if err != nil {
		return "", fmt.Errorf("Invalid YAML format")
	}
	return pretty_json, nil
}

func FormatJSON(rawJSON string) (string, error) {
	var prettyJSON bytes.Buffer

	// Try unmarshaling the JSON into a map or slice to check if it's valid.
	var jsonObj interface{}
	if err := json.Unmarshal([]byte(rawJSON), &jsonObj); err != nil {
		return "", errors.New("Invalid JSON format")
	}

	// Marshal the object with indentation.
	if err := json.Indent(&prettyJSON, []byte(rawJSON), "", "  "); err != nil {
		return "", errors.New("Failed to format JSON")
	}

	return prettyJSON.String(), nil
}

func CopyToClipboard(s string) {
	clipboard.WriteAll(s)
}

// pasteFromClipboard pastes the clipboard content into the paragraph widget.
func PasteFromClipboard(p *widgets.Paragraph) {
	text, err := clipboard.ReadAll()
	if err != nil {
		log.Println("Failed to read from clipboard:", err)
		return
	}
	p.Text = text
}

func EncodeBase64(s string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(s))
	return encoded
}

func DecodeBase64(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "Invalid Base64 string", err
	}
	return string(decoded), nil
}
