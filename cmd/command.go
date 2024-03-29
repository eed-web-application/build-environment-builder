package cmd

import (
	"errors"
	"fmt"

	"github.com/eed-web-application/build-environment-builder/cbs"
	"github.com/eed-web-application/build-environment-builder/utility"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(commandCMD)
	commandCMD.PersistentFlags().StringP("label", "l", "", "specify the endpoint label")

	commandCMD.AddCommand(commandCreateNew)
	commandCMD.AddCommand(listAllCommand)
	commandCMD.AddCommand(showCommand)
	commandCMD.AddCommand(updateCommand)
}

var commandCMD = &cobra.Command{
	Use:   "command",
	Short: "Command temaplte management",
	Long:  `Manage the command of the build ssytem`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

var commandCreateNew = &cobra.Command{
	Use:   "create",
	Short: "Create a new command tempalte",
	Long: `A comand tempalte is a set of commands that can be executed in a build system
	it consist in the following yaml representation:
	name: copy
	description: copy file or directory from source to destination
	parameters:
	- name: source
		description: source file or directory path
	- name: destination_dir
		description: destination directory path
	commandExecutionsLayers:
	- engine: shell
		architecture: 
		- linux
		operatingSystem:
		- ubuntu
		- redhat 
		executionCommands: 
		- cp -r ${source} ${destination_dir}
	`,

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
		path, err := utility.GetTemplate("new_command.yaml")
		if err != nil {
			return err
		}
		err = utility.EditFile(path)
		if err != nil {
			return err
		}
		var new_command cbs.NewCommandTemplateDTO
		err = utility.Deserialize(path, &new_command)
		if err != nil {
			return err
		}
		id, err := cbs.CreateNewCommandTemplate(endpoint, &new_command)
		if err != nil {
			return err
		}
		fmt.Printf("The new command template has been created with id: %s\n", *id)
		return nil
	},
}

var listAllCommand = &cobra.Command{
	Use:   "list",
	Short: "List all command",
	Long:  `Return a full list of the command in the build system`,

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

		list, err := cbs.FindAllCommand(endpoint)
		if err != nil {
			return err
		}
		idColor := color.New(color.FgGreen, color.Bold)
		nameColor := color.New(color.FgYellow, color.Italic)
		for _, command := range *list {
			fmt.Printf(
				"[%s] %s\n",
				idColor.Sprint(*command.Id),
				nameColor.Sprint(*command.Name),
			)
		}
		return nil
	},
}

var showCommand = &cobra.Command{
	Use:   "show",
	Args:  cobra.ExactArgs(1),
	Short: "Show command using the id",
	Long:  `Return the full command information`,

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

		c, err := cbs.FindCommandById(endpoint, args[0])
		if err != nil {
			return err
		}
		yaml, err := yaml.Marshal(c)
		if err != nil {
			return err
		}
		cy, err := utility.GetColorizedYaml(string(yaml))
		if err != nil {
			return err
		}
		fmt.Println(*cy)
		return nil
	},
}

var updateCommand = &cobra.Command{
	Use:   "update",
	Args:  cobra.ExactArgs(1),
	Short: "Update a command the id",
	Long:  `Fetch current command and open the editor to update it`,

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

		c, err := cbs.FindCommandById(endpoint, args[0])
		if err != nil {
			return err
		}
		yaml, err := yaml.Marshal(c)
		if err != nil {
			return err
		}
		path, err := utility.CreateTempFile(yaml)
		if err != nil {
			return err
		}
		beforeEditTime, err := utility.GetFileModTime(path)
		if err != nil {
			return fmt.Errorf("failed to get file modification time: %v", err)
		}

		err = utility.EditFile(path)
		if err != nil {
			return err
		}
		// Check the file's modification time after the editor has closed
		afterEditTime, err := utility.GetFileModTime(path)
		if err != nil {
			return fmt.Errorf("failed to get file modification time: %v", err)
		}
		// Compare the modification times
		if beforeEditTime == afterEditTime {
			fmt.Println("The file has not been modified.")
			return nil
		}
		var updated_command cbs.UpdateCommandTemplateDTO
		err = utility.Deserialize(path, &updated_command)
		if err != nil {
			return err
		}
		err = cbs.UpdateCommandById(endpoint, args[0], &updated_command)
		if err != nil {
			return err
		}
		fmt.Println("The command has been updated")
		return nil
	},
}
