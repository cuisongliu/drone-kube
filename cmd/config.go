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
	"github.com/cuisongliu/drone-kube/config"
	"github.com/cuisongliu/drone-kube/tools"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config your kubernetes classpath for admin",
	Run: func(cmd *cobra.Command, args []string) {
		//cmd.Flags().StringVarP(&name, "name", "n", "", "person's name")
		//cmd.Flags().IntVarP(&age, "age", "a", 0, "person's age")
		config.Main()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	kubeServer := tools.Env("KUBE_SERVER", "PLUGIN_SERVER", "PLUGIN_KUBE_SERVER")
	kubeCa := tools.Env("KUBE_CA", "PLUGIN_CA", "PLUGIN_KUBE_CA")
	kubeAdmin := tools.Env("KUBE_ADMIN", "PLUGIN_ADMIN", "PLUGIN_KUBE_ADMIN")
	kubeAdminKey := tools.Env("KUBE_ADMIN_KEY", "PLUGIN_ADMIN_KEY", "PLUGIN_KUBE_ADMIN_KEY")

	configCmd.Flags().StringVarP(&config.KubeServer, "server", "s", kubeServer, "~/.kube/config  server")
	//certificate-authority ca.pem
	configCmd.Flags().StringVarP(&config.KubeCa, "ca", "c", kubeCa, "~/.kube/config certificate-authority-data")
	//client-certificate admin.pem
	configCmd.Flags().StringVarP(&config.KubeAdmin, "admin", "n", kubeAdmin, "~/.kube/config client-certificate-data")
	//client-key admin-key.pem
	configCmd.Flags().StringVarP(&config.KubeAdminKey, "admin-key", "k", kubeAdminKey, "~/.kube/config client-key-data")
}
