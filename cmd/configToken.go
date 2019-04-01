// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/cuisongliu/drone-kube/configToken"
	"github.com/cuisongliu/drone-kube/tools"

	"github.com/spf13/cobra"
)

// configTokenCmd represents the configToken command
var configTokenCmd = &cobra.Command{
	Use:   "configToken",
	Short: "config your kubernetes classpath for token",
	Run: func(cmd *cobra.Command, args []string) {
		configToken.Main()
	},
}

func init() {
	rootCmd.AddCommand(configTokenCmd)

	kubeServer := tools.Env("KUBE_SERVER", "PLUGIN_SERVER", "PLUGIN_KUBE_SERVER")
	kubeUser := tools.Env("KUBE_USER", "PLUGIN_USER", "PLUGIN_KUBE_USER")
	kubeToken := tools.Env("KUBE_TOKEN", "PLUGIN_TOKEN", "PLUGIN_KUBE_TOKEN")

	configTokenCmd.Flags().StringVarP(&configToken.KubeServer, "server", "s", kubeServer, "~/.kube/config  server")
	configTokenCmd.Flags().StringVarP(&configToken.KubeUser, "user", "u", kubeUser, "~/.kube/config  user")
	configTokenCmd.Flags().StringVarP(&configToken.KubeToken, "token", "t", kubeToken, "~/.kube/config  token")

}
