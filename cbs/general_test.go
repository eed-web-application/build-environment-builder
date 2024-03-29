package cbs_test

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/eed-web-application/build-environment-builder/cbs"
)

func GetTestHost() string {
	host := "http://cbs:8080"
	if envHost := os.Getenv("CBS_TEST_HOST"); envHost != "" {
		fmt.Printf("Using host from environment variable: %s\n", envHost)
		host = envHost
	}
	return host
}

func resetData(host string) {
	client, clientErr := cbs.GetClient(host)
	if clientErr != nil {
		log.Fatalf("Error creating client: %v", clientErr)
	}

	result, err := client.DeleteAllWithResponse(context.Background())
	if err != nil {
		log.Fatalf("error calling API: %v", err)
	}
	if result.JSON201 != nil && result.JSON201.ErrorCode != 0 {
		log.Fatalf("error calling API: %v", result.JSON201.ErrorMessage)
	} else if result.JSON201 == nil {
		log.Fatalf("error calling API: %v", "No response")
	}
}

// func deserialize(file string, object any) {
// 	var intermediateData map[string]interface{}
// 	yamlFile, err := os.ReadFile(file)
// 	if err != nil {
// 		log.Fatalf("Failed to read YAML file: %v", err)
// 	}
// 	err = yaml.Unmarshal(yamlFile, &intermediateData)
// 	if err != nil {
// 		log.Fatalf("Failed to parse YAML: %v", err)
// 	}
// 	// Step 2: Convert the intermediate map to JSON
// 	jsonData, err := json.Marshal(intermediateData)
// 	if err != nil {
// 		log.Fatalf("Failed to marshal intermediate data to JSON: %v", err)
// 	}

// 	// Step 3: Unmarshal JSON into the final Go struct
// 	if err := json.Unmarshal(jsonData, object); err != nil {
// 		log.Fatalf("Failed to unmarshal JSON into Go struct: %v", err)
// 	}
// }
