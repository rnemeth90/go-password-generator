/*
Copyright © 2022 Ryan Nemeth ryan@geekyryan.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/rnemeth90/generator"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-password-generator",
	Short: "go-password-generator - a simple app for generating passwords",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

var generateMemorablePasswordCmd = &cobra.Command{
	Use:     "memorable",
	Aliases: []string{"mem"},
	Short:   "Generates a memorable password",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p, e := generator.GenerateMemorablePassword()
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println(p)
	},
}

var generateRandomCharPasswordCmd = &cobra.Command{
	Use:     "random",
	Aliases: []string{"rand"},
	Short:   "Generates a memorable password",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p, e := generator.GenerateRandomString(16)
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println(p)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-password-generator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(generateMemorablePasswordCmd)
	rootCmd.AddCommand(generateRandomCharPasswordCmd)
}
