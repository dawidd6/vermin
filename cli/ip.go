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
package cli

import (
	"errors"
	"fmt"
	"github.com/mhewedy/vermin/vms"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Show IP address for a running VM",
	Long: `Show IP address for a running VM

Sometimes, the IP information being stale, so you might need use the --purge|-p flag

Examples:
$ vermin ip vm_11
To purge the IP cache:
$ vermin ip vm_05 -p
`,
	Run: func(cmd *cobra.Command, args []string) {
		vmName := args[0]
		purge, _ := cmd.Flags().GetBool("purge")

		ps, err := vms.IP(vmName, purge)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(ps)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("vm required")
		}
		return nil
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		list, _ := vms.List(false)
		var completions []string
		for _, comp := range list {
			if strings.HasPrefix(comp, toComplete) {
				completions = append(completions, comp)
			}
		}
		return completions, cobra.ShellCompDirectiveDefault
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	ipCmd.Flags().BoolP("purge", "p", false, "Purge the IP cache")
}
