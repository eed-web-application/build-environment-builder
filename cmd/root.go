// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// ----------------------------------------------------------------------------
// COBRA COMMAND
// ----------------------------------------------------------------------------

type RootFlagsType struct {
	Verbose string
}

var RootFlags = &RootFlagsType{
	Verbose: "debug",
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "eed build system",
	Short:        "EED build system - Command Line Tools.",
	Long:         `EED build system long descirption`,
	SilenceUsage: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		v, _ := cmd.Flags().GetString("verbose")
		v = strings.ToLower(v)
		setLogLevel(v)

	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func initConfig() {
	logrus.Debug("Configuration init")
	ReadConfiguration()
	setLogLevel(RootFlags.Verbose)
}

func init() {
	cobra.OnInitialize(initConfig)
	logrus.Debug("Comeplte configuration init")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&RootFlags.Verbose, "verbose", "v", "info", "Verbose level {error | warn | info | debug | trace}")
}

// setup the log level
func setLogLevel(l string) (logrus.Level, logrus.Level) {
	old := logrus.GetLevel()
	l = strings.ToLower(l)
	switch l {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.Fatal("Wrong level. Use trace, debug, info, warn, error.")
	}
	new := logrus.GetLevel()
	return old, new
}

// ----------------------------------------------------------------------------
// BUSINESS LOGIN
// ----------------------------------------------------------------------------

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetOutput(os.Stderr)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
