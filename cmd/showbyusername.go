/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"log"

	"github.com/G0tYou/user-service-cli/helper"
	userPB "github.com/G0tYou/user-service/proto"
	"github.com/spf13/cobra"
)

// showbyusernameCmd represents the showbyusername command
var showbyusernameCmd = &cobra.Command{
	Use:   "showbyusername",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		showbyusername(args)
	},
}

func init() {
	rootCmd.AddCommand(showbyusernameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showbyusernameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showbyusernameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func showbyusername(args []string) {
	client := NewClient()

	// Parsing the argument
	file := args[0]
	user, err := helper.ParseFile(file)
	if err != nil {
		log.Fatalf("could not parse the argument provided %v", err)
	}

	// Call ShowUser rpc from grpc client
	res, err := client.ShowUserByUsername(context.Background(), &userPB.User{Username: string(user.Username)})
	if err != nil {
		log.Fatalf("could not show the user with username = %s %v", user.Username, err)
	}
	log.Println(res.User)
}
