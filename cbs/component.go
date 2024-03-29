package cbs

import (
	"context"
	"fmt"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/sirupsen/logrus"
)

// Replace "path/to/cbsapi" with the actual import path of the "cbsapi" package

func getClient(host string) (*ClientWithResponses, error) {
	apiKeyProvider, apiKeyProviderErr := securityprovider.NewSecurityProviderApiKey("header", "X-API-Key", "MY_API_KEY")
	if apiKeyProviderErr != nil {
		logrus.Error(fmt.Printf("error setting the security provider: %v", apiKeyProviderErr))
		return nil, apiKeyProviderErr
	}
	client, clientErr := NewClientWithResponses(host, WithRequestEditorFn(apiKeyProvider.Intercept))
	if clientErr != nil {
		logrus.Error(fmt.Printf("error creating client: %v", clientErr))
		return nil, clientErr
	}
	return client, nil
}

// FindAllComponent is the function to get all the components
func FindAllComponent(host string) (*[]ComponentSummaryDTO, error) {
	client, clientErr := getClient(host)
	if clientErr != nil {
		return nil, clientErr
	}

	result, err := client.ListAllComponentWithResponse(context.Background())
	if err != nil {
		logrus.Error(fmt.Printf("error calling API: %v", err))
		return nil, err
	}
	return result.JSON200.Payload, nil
}

// CreateNewComponent create a new component  and return the id
func CreateNewComponent(host string, component NewComponentDTO) (*string, error) {
	client, clientErr := getClient(host)
	if clientErr != nil {
		return nil, clientErr
	}

	result, err := client.CreateWithResponse(context.Background(), component)
	if err != nil {
		logrus.Error(fmt.Printf("error calling API: %v", err))
		return nil, err
	}
	if result.JSON201 != nil && result.JSON201.ErrorCode != 0 {
		return nil, fmt.Errorf("error calling API: %v", result.JSON201.ErrorMessage)
	} else if result.JSON201 == nil {
		return nil, fmt.Errorf("error calling API: %v", "No response")
	}
	return result.JSON201.Payload, nil
}

// CreateNewCommandTemplate create a new command template and return the id
func CreateNewCommandTemplate(host string, new_command NewCommandTemplateDTO) (*string, error) {
	client, clientErr := getClient(host)
	if clientErr != nil {
		return nil, clientErr
	}

	result, err := client.Create1WithResponse(context.Background(), new_command)
	if err != nil {
		logrus.Error(fmt.Printf("error calling API: %v", err))
		return nil, err
	}
	if result.JSON201 != nil && result.JSON201.ErrorCode != 0 {
		return nil, fmt.Errorf("error calling API: %v", result.JSON201.ErrorMessage)
	} else if result.JSON201 == nil {
		return nil, fmt.Errorf("error calling API: %v", "No response")
	}
	return result.JSON201.Payload, nil
}

// DeleteCommandTemplate delete a command template by id
func DeleteCommandTemplate(host string, id string) error {
	client, clientErr := getClient(host)
	if clientErr != nil {
		return clientErr
	}

	result, err := client.DeleteById1WithResponse(context.Background(), id)
	if err != nil {
		logrus.Error(fmt.Printf("error calling API: %v", err))
		return err
	}
	if result.JSON200 != nil && result.JSON200.ErrorCode != 0 {
		return fmt.Errorf("error calling API: %v", result.JSON200.ErrorMessage)
	} else if result.JSON200 == nil {
		return fmt.Errorf("error calling API: %v", "No response")
	}
	return nil
}
