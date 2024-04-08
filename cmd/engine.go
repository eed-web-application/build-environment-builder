package cmd

import (
	"errors"
	"fmt"

	"github.com/eed-web-application/build-environment-builder/cbs"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(engineCMD)
	engineCMD.PersistentFlags().StringP("label", "l", "", "specify the endpoint label")

	engineCMD.AddCommand(listAllEngine)
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
