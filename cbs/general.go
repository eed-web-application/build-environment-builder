package cbs

import (
	"encoding/json"
	"fmt"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/sirupsen/logrus"
)

// GetClient is the function to get the client
func GetClient(host string) (*ClientWithResponses, error) {
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

// DecodeReponse is the function to decode the response
func DecodeReponse(bytes []byte) (*map[string]interface{}, error) {
	var deserialized map[string]interface{}
	if err := json.Unmarshal(bytes, &deserialized); err != nil {
		return nil, err
	}
	return &deserialized, nil
}
