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
	"strconv"

	"github.com/micro/go-micro/metadata"

	userPB "github.com/G0tYou/user-service/proto"
	"github.com/spf13/cobra"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		destroy(args)
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// destroyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// destroyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func destroy(args []string) {
	client := NewClient()

	// Parsing the argument
	id, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		log.Fatalf("could not parse the argument provided %v", err)
	}

	token := args[1]
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	// Call DestroyUser rpc from grpc client
	res, err := client.DestroyUser(ctx, &userPB.User{Id: int32(id)})
	if err != nil {
		log.Printf("could not destroy the user with id = %d %v", id, err)
	}
	log.Println(res.User)
}
