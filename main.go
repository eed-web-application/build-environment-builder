package main

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"gopkg.in/yaml.v3"
)

type Config struct {
	BaseImage string   `yaml:"base_image"`
	Commands  []string `yaml:"commands"`
	// Add other fields as needed
}

var wg sync.WaitGroup

// Function to parse YAML file
func parseYAML(yamlFile string) (*Config, error) {
	var config Config
	data, err := os.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func buildImage(cli *client.Client, config *Config) error {

	wg.Add(1)
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	// Dynamically create a Dockerfile from the config
	dockerfile := fmt.Sprintf("FROM %s\n", config.BaseImage)
	for _, cmd := range config.Commands {
		dockerfile += fmt.Sprintf("RUN %s\n", cmd)
	}

	// Add Dockerfile to the tar archive
	if err := tw.WriteHeader(&tar.Header{
		Name: "Dockerfile",
		Size: int64(len(dockerfile)),
	}); err != nil {
		return err
	}

	if _, err := tw.Write([]byte(dockerfile)); err != nil {
		return err
	}

	if err := tw.Close(); err != nil {
		return err
	}

	// Use the tar archive as the build context
	buildContext := bytes.NewReader(buf.Bytes())

	opts := types.ImageBuildOptions{
		Tags:       []string{"your-image-name:latest"},
		Remove:     true,
		Dockerfile: "Dockerfile",
	}

	response, err := cli.ImageBuild(context.Background(), buildContext, opts)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Process build output asynchronously
	go func() {
		defer wg.Done()
		// Create an io.Writer to capture the output, for example, os.Stdout or a buffer
		outputWriter := os.Stdout // You can use any io.Writer here

		// Decode the response to get the stream of build messages
		var message json.RawMessage
		decoder := json.NewDecoder(response.Body)

		for {
			if err := decoder.Decode(&message); err != nil {
				if err == io.EOF {
					break
				}
				fmt.Fprintf(outputWriter, "Error reading JSON message: %v\n", err)
				return
			}

			var jsonMessage map[string]interface{}
			if err := json.Unmarshal(message, &jsonMessage); err != nil {
				fmt.Fprintf(outputWriter, "Error unmarshaling JSON message: %v\n", err)
				continue
			}

			if stream, ok := jsonMessage["stream"].(string); ok {
				fmt.Fprintf(outputWriter, stream)
			} else if errMsg, ok := jsonMessage["errorDetail"].(map[string]interface{}); ok && errMsg["message"] != nil {
				fmt.Fprintf(outputWriter, "Error: %s\n", errMsg["message"].(string))
				return
			}
		}
	}()

	// Your main function can continue executing other tasks here
	// Remember to handle synchronization properly if necessary
	wg.Wait() // This will block until wg.Done() is called.
	return nil
}

func createDockerClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
		return nil, err
	}
	fmt.Println("Successfully connected to Docker daemon")
	return cli, nil
}

func main() {
	// Parse command-line arguments to get the YAML file path
	// For simplicity, error handling is omitted here
	yamlFile := "test-build.yaml"

	config, err := parseYAML(yamlFile)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %v", err)
	}

	cli, err := createDockerClient()
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	if err := buildImage(cli, config); err != nil {
		log.Fatalf("Error building Docker image: %v", err)
	}

}
