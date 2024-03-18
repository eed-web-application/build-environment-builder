package main

import "github.com/eed-web-application/build-environment-builder/cmd"

func main() {
	cmd.Execute()
}

// func main() {
// 	// Parse command-line arguments to get the YAML file path
// 	// For simplicity, error handling is omitted here
// 	yamlFile := "test-build.yaml"

// 	config, err := parseYAML(yamlFile)
// 	if err != nil {
// 		log.Fatalf("Error parsing YAML file: %v", err)
// 	}

// 	cli, err := createDockerClient()
// 	if err != nil {
// 		log.Fatalf("Error creating Docker client: %v", err)
// 	}

// 	if err := buildImage(cli, config); err != nil {
// 		log.Fatalf("Error building Docker image: %v", err)
// 	}

// }
