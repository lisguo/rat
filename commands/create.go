/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	lib "github.com/lisguo/resume-gen/lib"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resume using a particular template",
	Long: `
	resume-gen is a resume generator CLI tool that uses LaTeX resume templates to  
	create professional resumes.
	
	To start try resume-gen create sample_1 -f sample_configs/tech.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		// Create file
		f, err := os.Create("test.tex")
		if err != nil {
			panic(err)
		}

		lib.WriteHeader(f)

		// Write Contact
		name := viper.GetString("personal.name")
		address := viper.GetString("personal.address")
		phone := viper.GetString("personal.phone")
		email := viper.GetString("personal.email")
		lib.WriteContactInfo(f, name, address, phone, email)
		lib.WriteBeginDocument(f)

		// Write Education
		lib.WriteBeginSection(f, "Education")

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
