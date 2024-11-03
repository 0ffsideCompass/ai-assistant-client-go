package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/models"
)

const (
	generateEndpoint = "/generate"
)

func (c *Client) Generate(prompt, languge string, contentType models.GeneratedContentType) (*models.GenerateRequest, error) {
	payload := models.GenerateRequest{
		Prompt: prompt,
		Lang:   languge,
		Type:   contentType,
	}

	responseData, err := c.post(generateEndpoint, payload)
	if err != nil {
		return nil, fmt.Errorf("error posting league: %w", err)
	}

	var data models.GenerateResponse
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return &data, nil
}
