/*
Copyright 2020 The blcli Authors All rights reserved.
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
	"errors"
	"fmt"
	"os"

	"./printer"
	"github.com/spf13/cobra"
)

// CreateOptions sets up the create options command and subcommands
func CreateOptions() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-options <bitlaunch, digitalocean, vultr, or linode>",
		Short:   "View images, sizes, and options available for a host when creating a new server.",
		Long:    ``,
		Aliases: []string{"o"},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("please provide a host name: bitlaunch, digitalocean, vultr, or linode")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			hid, err := hostID(id)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			server, err := client.CreateOptions.Show(hid)
			if err != nil {
				fmt.Printf("Error getting server : %v\n", err)
				os.Exit(1)
			}

			printer.Output(server)
		},
	}
	return cmd
}
