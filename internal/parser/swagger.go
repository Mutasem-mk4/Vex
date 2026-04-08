package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// SwaggerData is the simplified struct structure for extracting paths only
type SwaggerData struct {
	Paths map[string]map[string]interface{} `json:"paths"`
}

// ParseSwagger ingests a Swagger/OpenAPI v2/v3 JSON file and returns an automated string slice of endpoints
func ParseSwagger(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read swagger file: %w", err)
	}

	var swag SwaggerData
	if err := json.Unmarshal(data, &swag); err != nil {
		return nil, fmt.Errorf("failed to parse json logic: %w", err)
	}

	var endpoints []string
	
	// Iterate gracefully over all paths and http methods (get, post, put, etc.)
	for path, methods := range swag.Paths {
		for method := range methods {
			// e.g. "GET /api/v1/users/{id}"
			ep := fmt.Sprintf("%s %s", strings.ToUpper(method), path)
			endpoints = append(endpoints, ep)
		}
	}

	return endpoints, nil
}
