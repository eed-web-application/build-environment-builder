package cbs

import (
	"bytes"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

// FetchAllEngines fetch all engines
func FetchAllEngines(host string) (*[]string, error) {
	client, clientErr := GetClient(host)
	if clientErr != nil {
		return nil, clientErr
	}

	result, err := client.FindAllEngineNamesWithResponse(context.Background())
	if err != nil {
		logrus.Error(fmt.Printf("error calling API: %v", err))
		return nil, err
	}
	if result.JSON200 != nil && result.JSON200.ErrorCode != 0 {
		return nil, fmt.Errorf("error calling API: %v", result.JSON200.ErrorMessage)
	} else if result.JSON200 == nil {
		if json, err := DecodeReponse(result.Body); err == nil {
			return nil, fmt.Errorf("error calling API: %v", (*json)["errorMessage"])
		} else {
			return nil, fmt.Errorf("error calling API: %v", "No response")
		}
	}
	return result.JSON200.Payload, nil
}

// GenerateComponentArtifact fetch all engines
func GenerateComponentArtifact(host string, param *GenerateComponentArtifactParams) (*bytes.Reader, error) {
	client, clientErr := GetClient(host)
	if clientErr != nil {
		return nil, clientErr
	}

	result, err := client.GenerateComponentArtifactWithResponse(context.Background(), param)
	if err != nil {
		logrus.Error(fmt.Printf("error calling API: %v", err))
		return nil, err
	}
	if result.Body != nil {
		return bytes.NewReader(result.Body), nil
	} else if result.JSON200 == nil {
		if json, err := DecodeReponse(result.Body); err == nil {
			return nil, fmt.Errorf("error calling API: %v", (*json)["errorMessage"])
		} else {
			return nil, fmt.Errorf("error calling API: %v", "No response")
		}
	}
	return nil, nil
}
