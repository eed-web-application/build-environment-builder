package cbs

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

// CreateNewCommandTemplate create a new command template and return the id
func CreateNewCommandTemplate(host string, new_command *NewCommandTemplateDTO) (*string, error) {
	client, clientErr := GetClient(host)
	if clientErr != nil {
		return nil, clientErr
	}

	result, err := client.CreateCommandWithResponse(context.Background(), *new_command)
	if err != nil {
		logrus.Error(fmt.Printf("error calling API: %v", err))
		return nil, err
	}
	if result.JSON201 != nil && result.JSON201.ErrorCode != 0 {
		return nil, fmt.Errorf("error calling API: %v", result.JSON201.ErrorMessage)
	} else if result.JSON201 == nil {
		if json, err := DecodeReponse(result.Body); err == nil {
			return nil, fmt.Errorf("error calling API: %v", (*json)["errorMessage"])
		} else {
			return nil, fmt.Errorf("error calling API: %v", "No response")
		}
	}
	return result.JSON201.Payload, nil
}

// DeleteCommandTemplate delete a command template by id
func DeleteCommandTemplate(host string, id string) error {
	client, clientErr := GetClient(host)
	if clientErr != nil {
		return clientErr
	}

	result, err := client.DeleteCommandByIdWithResponse(context.Background(), id)
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

// FindCommandById is the function to get all the command templates
func FindCommandById(host string, id string) (*CommandTemplateDTO, error) {
	client, clientErr := GetClient(host)
	if clientErr != nil {
		return nil, clientErr
	}

	result, err := client.FindCommandByIdWithResponse(context.Background(), id)
	if err != nil {
		logrus.Error(fmt.Printf("error calling API: %v", err))
		return nil, err
	}
	if result.JSON200 != nil && result.JSON200.ErrorCode != 0 {
		return nil, fmt.Errorf("error calling API: %v", result.JSON200.ErrorMessage)
	} else if result.JSON200 == nil {
		return nil, fmt.Errorf("error calling API: %v", "No response")
	}
	return result.JSON200.Payload, nil
}

// UpdateCommandById update a command template by id
func UpdateCommandById(host string, id string, update_command_dto *UpdateCommandTemplateDTO) error {
	client, clientErr := GetClient(host)
	if clientErr != nil {
		return clientErr
	}

	result, err := client.UpdateCommandByIdWithResponse(context.Background(), id, *update_command_dto)
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

// FindAllCommand is the function to get all the command templates
func FindAllCommand(host string) (*[]CommandTemplateSummaryDTO, error) {
	client, clientErr := GetClient(host)
	if clientErr != nil {
		return nil, clientErr
	}

	result, err := client.ListAllCommandWithResponse(context.Background())
	if err != nil {
		logrus.Error(fmt.Printf("error calling API: %v", err))
		return nil, err
	}
	if result.JSON200 != nil && result.JSON200.ErrorCode != 0 {
		return nil, fmt.Errorf("error calling API: %v", result.JSON200.ErrorMessage)
	} else if result.JSON200 == nil {
		return nil, fmt.Errorf("error calling API: %v", "No response")
	}
	return result.JSON200.Payload, nil
}
