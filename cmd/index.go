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
	userPB "github.com/G0tYou/user-service/proto"
	"github.com/G0tYou/user-service-cli/cmd"
	"log"
	"google.golang.org/grpc"
	"github.com/spf13/cobra"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		index()
	},
}

func index() {
	// Connect to grpc client
	conn, err := grpc.Dial(cmd.Address, grpc.WithInsecure())
	if err != nil{
		log.Fatalf("did not connect: #{err}")
	}
	defer conn.Close()
	client := userPB.NewUserServiceClient(conn)

	// Call indexUser rpc from grpc client
	res, err:= client.IndexUser(context.Background(), &userPB.IndexUsersRequest{})
	if err != nil{
		log.Fatalf("could not index users #{err}")
	}
	for _, user := range res.Users {
		log.Println(user)
	}
}

func init() {
	rootCmd.AddCommand(indexCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// indexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// indexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func index() {
	client := NewClient()

	// Call IndexUsers rpc from grpc client
	res, err := client.IndexUsers(context.Background(), &userPB.IndexUsersRequest{})
	if err != nil {
		log.Fatalf("could not index users %v", err)
	}
	for _, user := range res.Users {
		log.Println(user)
	}
}
