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

	"github.com/sahilm/fuzzy"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [SEARCH_QUERY]",
	Short: "Provide a search query to return all docker images that contain the search query",
	Long: `Search command to search through the images in the remote repository.
	Provide a search term as an argument.`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		// Get repository url
		url := viper.GetString("repo_url")
		fmt.Println("Repo URL:", url)

		searchQuery := args[0]
		fmt.Println("Searching for:", searchQuery, "in", url)

		catalogStr := "https://" + url + "/v2/_catalog"

		// Create a Get request to retrieve the docker repositories in the registry
		resp, err := http.Get(catalogStr)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		repositoriesInterface := result["repositories"]

		// String Matching
		// Uses Fuzzy Matching to search within the list of image repositories
		// Highlights the matched word with the repository
		const bold = "\033[36m%s\033[0m"
		data := []string{}

		for _, value := range repositoriesInterface.([]interface{}) {
			data = append(data, value.(string))
		}

		matches := fuzzy.Find(searchQuery, data)

		fmt.Println()
		if len(matches) == 0 {
			fmt.Println("No matches found.")
		} else {
			fmt.Println("Match Results: ")
		}

		for _, match := range matches {
			for i := 0; i < len(match.Str); i++ {
				if contains(i, match.MatchedIndexes) {
					fmt.Print(fmt.Sprintf(bold, string(match.Str[i])))
				} else {
					fmt.Print(string(match.Str[i]))
				}
			}
			fmt.Println()
		}

	},
}

func contains(needle int, haystack []int) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
