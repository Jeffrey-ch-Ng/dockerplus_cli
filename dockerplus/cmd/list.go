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
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all docker images in the docker registry",
	Long:  `List all docker image repositories that exist in the registry`,
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Repo URL: ", viper.GetString("repo_url"))

		if viper.Get("repo_url") == nil {
			fmt.Println("Enter `dockerplus set <server_name>` to set the server name as an environment variable before you move on.")
			return
		}

		catalogStr := "https://" + viper.Get("repo_url").(string) + "/v2/_catalog"

		resp, err := http.Get(catalogStr)
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)

		repositoriesInterface := result["repositories"]

		// Print the docker images in the docker registry
		for key, value := range repositoriesInterface.([]interface{}) {
			fmt.Println(key, ":", value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
