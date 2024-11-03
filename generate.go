package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/models"
)

const (
	generateEndpoint = "/generate"
)

func (c *Client) Generate(prompt, language string, contentType models.GeneratedContentType) (*models.GenerateResponse, error) {
	payload := models.GenerateRequest{
		Prompt: prompt,
		Lang:   language,
		Type:   contentType,
	}

	responseData, err := c.post(generateEndpoint, payload)
	if err != nil {
		return nil, fmt.Errorf("error posting generate request: %w", err)
	}

	var data models.GenerateResponse
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return &data, nil
}
