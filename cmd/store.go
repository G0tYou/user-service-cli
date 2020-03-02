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
	"github.com/G0tYou/user-service-cli/helper"
	"github.com/SleepingNext/product-service-cli/cmd"
	"log"
	pb"github.com/G0tYou/user-service/proto"
	"google.golang.org/grpc"
	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		store(args)
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func store(args []string) {
	// Connect to grpc client
	conn, err := grpc.Dial(cmd.Address, grpc.WithInsecure())
	if err != nil{
		log.Fatalf("did not connect: #{err}")
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	file := args[0]
	user, err := helper.ParseFile(file)
	if err != nil {
		log.Fatalf("could not parse file: #{err}")
	}

	//call StoreUser rpc from grpc client
	res, err := client.StoreUser(context.Background(), user)
	if err != nil{
		log.Fatalf("could not store the product in file = #{file} #{err}")
	}
	log.Println(res.User)
}