package cmd

import (
	"errors"
	"fmt"

	"github.com/eed-web-application/build-environment-builder/cbs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ----------------------------------------------------------------------------
// INIT
// ----------------------------------------------------------------------------

func init() {
	rootCmd.AddCommand(componentCmd)
	componentCmd.PersistentFlags().StringP("label", "l", "", "specify the endpoint label")

	componentCmd.AddCommand(componentFindAllCmd)

	// godivaRuoliServiziCmd.Flags().String("anagId", "", "uuid oppure id anagrafico del servizio")
	// godivaRuoliServiziCmd.Flags().Int("nodoId", 0, "id nodo dell'applicazione")
	// godivaRuoliServiziCmd.Flags().Bool("all", false, "mostra i ruoli per tutte le anagrafiche di tipo servizio")
}

var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Component Management",
	Long:  `Manage the component of the build ssytem`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

var componentFindAllCmd = &cobra.Command{
	Use:   "find-all",
	Short: "Find all compoenent",
	Long:  `Reaturn all compoenet in the system`,

	RunE: func(cmd *cobra.Command, args []string) error {
		var foundErr error
		var found_component *[]cbs.ComponentDTO

		label_endpoint, _ := cmd.Flags().GetString("label")
		if label_endpoint == "" {
			return errors.New("the label parameter is mandatory")
		}
		endpoint, ok := Configuration.Endpoints[label_endpoint]
		if !ok {
			return errors.New("the label is not configured")
		}
		logrus.Debug("Use endpoint: ", endpoint)

		if found_component, foundErr = cbs.FindAllComponent(endpoint); foundErr != nil {
			return foundErr
		}
		for index, value := range *found_component {
			fmt.Printf("Index: %d, Value: %d\n", index, value.Name)
		}
		return nil
	},
}
