package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/eed-web-application/build-environment-builder/cbs"
	"github.com/eed-web-application/build-environment-builder/utility"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(engineCMD)
	engineCMD.PersistentFlags().StringP("label", "l", "", "specify the endpoint label")

	engineCMD.AddCommand(listAllEngine)
	engineCMD.AddCommand(generateCMD)
	generateCMD.Flags().StringP("engine", "e", "", "Specify the engine to use for the generation of the artifact")
	generateCMD.Flags().StringSlice("component", nil, "Component list for which we want to generate the artifact")
	generateCMD.Flags().StringSlice("epar", nil, "All the egine paramters to use for the generation of the artifact")
}

var engineCMD = &cobra.Command{
	Use:   "engine",
	Short: "Generate artifact from specific engines for different deployment systems",
	Long: `Generate artifact from specific engines for different deployment systems
	ad exmplae of engine is docker, ansible, etc
	`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

var listAllEngine = &cobra.Command{
	Use:   "list",
	Short: "Show all available engines",
	Long:  `Show all available engines managed by the backend`,

	RunE: func(cmd *cobra.Command, args []string) error {
		label_endpoint, _ := cmd.Flags().GetString("label")
		if label_endpoint == "" {
			return errors.New("the label parameter is mandatory")
		}
		endpoint, ok := Configuration.Endpoints[label_endpoint]
		if !ok {
			return errors.New("the label is not configured")
		}
		engine_list, err := cbs.FetchAllEngines(endpoint)
		if err != nil {
			return err
		}
		// ok we can show the result
		for _, engine := range *engine_list {
			fmt.Println(engine)
		}
		return nil
	},
}

//
// type GenerateComponentArtifactParams struct {
// 	// EngineName IS the engine to use represented by his name
// 	EngineName string `form:"engineName" json:"engineName"`

// 	// ComponentId Is the list of the component id for wich the artifact should be generated
// 	ComponentId []string `form:"componentId" json:"componentId"`

//		// AllRequestParams is the build specs to use for the generation of the artifact
//		AllRequestParams map[string]string `form:"allRequestParams" json:"allRequestParams"`
//	}
var generateCMD = &cobra.Command{
	Use:   "generate",
	Short: "Generate artifact starting from a specific engine",
	Long:  `Generate artifact starting from a specific engine`,

	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var engine string
		var components []string
		var key_value []string
		var artifact *bytes.Reader

		engine_param := make(map[string]string)

		label_endpoint, _ := cmd.Flags().GetString("label")
		if label_endpoint == "" {
			return errors.New("the label parameter is mandatory")
		}
		endpoint, ok := Configuration.Endpoints[label_endpoint]
		if !ok {
			return errors.New("the label is not configured")
		}
		if engine, err = cmd.Flags().GetString("engine"); err != nil {
			return err
		}
		if components, err = cmd.Flags().GetStringSlice("component"); err != nil {
			return err
		}
		if key_value, err = cmd.Flags().GetStringSlice("epar"); err != nil {
			return err
		}
		// ok we can show the result
		for _, pair := range key_value {
			parts := strings.SplitN(pair, "=", 2)
			if len(parts) == 2 {
				engine_param[parts[0]] = parts[1]
			} else {
				return errors.New("invalid key-value pair")
			}
		}

		if artifact, err = cbs.GenerateComponentArtifact(
			endpoint,
			&cbs.GenerateComponentArtifactParams{
				EngineName:       engine,
				ComponentId:      components,
				AllRequestParams: engine_param,
			},
		); err != nil {
			return err
		}
		// send file on the standard output
		return utility.WriteToStdOut(artifact)
	},
}
