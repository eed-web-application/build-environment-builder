package cmd

import (
	"errors"
	"fmt"

	"github.com/eed-web-application/build-environment-builder/cbs"
	"github.com/eed-web-application/build-environment-builder/utility"
	"github.com/fatih/color"
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
	componentCmd.AddCommand(createComponentCMD)

}

var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Component Management",
	Long:  `Manage the component of the build ssytem`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

var componentFindAllCmd = &cobra.Command{
	Use:   "list",
	Short: "List all compoenent",
	Long:  `Reaturn all component in the system`,

	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var found_component *[]cbs.ComponentSummaryDTO

		label_endpoint, _ := cmd.Flags().GetString("label")
		if label_endpoint == "" {
			return errors.New("the label parameter is mandatory")
		}
		endpoint, ok := Configuration.Endpoints[label_endpoint]
		if !ok {
			return errors.New("the label is not configured")
		}
		logrus.Debug("Use endpoint: ", endpoint)

		if found_component, err = cbs.FindAllComponent(endpoint); err != nil {
			return err
		}
		if len(*found_component) != 0 {
			idColor := color.New(color.FgGreen, color.Bold)
			nameColor := color.New(color.FgYellow, color.Italic)

			for _, c := range *found_component {
				fmt.Printf(
					"[%s] %s\n",
					idColor.Sprint(*c.Id),
					nameColor.Sprint(*c.Name),
				)
			}
		} else {
			fmt.Println("No component found")
		}
		return nil
	},
}

var createComponentCMD = &cobra.Command{
	Use:   "create",
	Short: "Create new component",
	Long:  `Create a new component in the system and return his id`,

	RunE: func(cmd *cobra.Command, args []string) error {
		label_endpoint, _ := cmd.Flags().GetString("label")
		if label_endpoint == "" {
			return errors.New("the label parameter is mandatory")
		}
		endpoint, ok := Configuration.Endpoints[label_endpoint]
		if !ok {
			return errors.New("the label is not configured")
		}
		logrus.Debug("Use endpoint: ", endpoint)
		path, err := utility.GetTemplate("new_component.yaml")
		if err != nil {
			return err
		}
		err = utility.EditFile(path)
		if err != nil {
			return err
		}
		var new_command cbs.NewComponentDTO
		err = utility.Deserialize(path, &new_command)
		if err != nil {
			return err
		}
		id, err := cbs.CreateNewComponent(endpoint, &new_command)
		if err != nil {
			return err
		}
		fmt.Printf("The new command template has been created with id: %s\n", *id)
		return nil
	},
}
