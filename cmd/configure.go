package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// CliConfig è un raccoglitore delle diverse tipologie di configurazioni
type CliConfig struct {
	LogLevel  string            `yaml:"defaultLogLevel"`
	Endpoints map[string]string `yaml:"enpoints"` // Add your map here
}

// Configurazioni contiene tutte le parametrizzazioni dell'applicazione
var Configuration = &CliConfig{
	LogLevel: "debug",
}

func init() {
	// add configure command to root command
	rootCmd.AddCommand(configure)

	// add configureApiEndpoint command to configure command
	configure.AddCommand(configureApiEndpoint)
	configureApiEndpoint.PersistentFlags().StringP("label", "l", "", "Specify the label to configure")
	configureApiEndpoint.PersistentFlags().StringP("url", "u", "", "Specify the url to configure")
}

// configureCmd represents the main configure command
var configure = &cobra.Command{
	Use:   "configure",
	Short: "Configure client",
	Long:  `Update the cbs client configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debug("Start configuration")
	},
}

// configureApiEndpoint permit to create cbs api endpoint associated to a label
var configureApiEndpoint = &cobra.Command{
	Use:   "endpoint",
	Short: "Configure an api endpoint",
	Long: `Configure or update an api endpoint associated with a label.
			Example: cbs configure endpoint --label=dev --url=http://localhost:8080/api/v1/`,
	Run: func(cmd *cobra.Command, args []string) {
		label, err := cmd.Flags().GetString("label")
		if err != nil || label == "" {
			logrus.Error("Missing or empty 'label' parameter")
			return
		}

		url, err := cmd.Flags().GetString("url")
		if err != nil {
			logrus.Error("Missing or empty 'url' parameter")
			return
		} else if url == "" {
			logrus.Error(fmt.Sprintf("Remove label %s from configuration", label))
			delete(Configuration.Endpoints, label)
		} else {
			logrus.Debug("configure endpoint")
			if Configuration.Endpoints == nil {
				Configuration.Endpoints = make(map[string]string)
			}
			Configuration.Endpoints[label] = url
		}

		SaveConfiguration()
	},
}

// ConfigPath ?
func ConfigPath() string {
	var configdir string
	if configdir != "" {
		return configdir
	}
	path := os.Getenv("CBS_CONFIG_DIR")
	if path == "" {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = home + "/.cbs"
		path = filepath.Clean(path)

	}
	return path
}

// ReadConfiguration
func ReadConfiguration() error {
	settingsFile := ConfigPath() + "/settings.yaml"
	settingsFile = filepath.Clean(settingsFile)

	logrus.Infof("Read configuration from --> %s", settingsFile)

	yamlFile, err := os.ReadFile(settingsFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// Create the config directory if it doesn't exist
			err := os.MkdirAll(filepath.Dir(settingsFile), 0755)
			if err != nil {
				log.Fatalf("os.MkdirAll: %v", err)
				return err
			}
			// Handle the case when the configuration file is not found.
			// For example, you can create a new default configuration file.
			// return err
			if err = SaveConfiguration(); err != nil {
				log.Fatalf("SaveConfiguration: %v", err)
				return err
			}
		} else {
			log.Fatalf("os.ReadFile: %v", err)
			return err
		}
	}

	err = yaml.Unmarshal(yamlFile, Configuration)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}
	return nil
}

// SaveConfiguration
func SaveConfiguration() error {
	settingsFile := ConfigPath() + "/settings.yaml"
	b, err := yaml.Marshal(Configuration)
	if err != nil {
		log.Fatalf("Marshal: %v", err)
		return err
	}
	err = ioutil.WriteFile(settingsFile, b, 0644)
	if err != nil {
		log.Fatalf("ioutil.WriteFile: %v", err)
		return err
	}
	return nil
}
