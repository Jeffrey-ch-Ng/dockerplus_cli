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
package main

import (
	"log"

	"github.com/Jeffrey-ch-Ng/dockerplus_cli/dockerplus/cmd"

	"github.com/spf13/viper"
)

func main() {

	// Read the config.yml file to retrieve the docker regsitry url
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath("$GOPATH/src/github.com/Jeffrey-ch-Ng/dockerplus_cli/dockerplus/")

	// viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	cmd.Execute()
}
