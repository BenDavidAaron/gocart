/*
Copyright © 2021 Ben Aaron <ben@betadeltaalpha.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	gocart "github.com/BenDavidAaron/gocart/internal"
	"github.com/spf13/cobra"
)

// configDelCmd represents the configDel command
var configDelCmd = &cobra.Command{
	Use:   "config",
	Short: "Delete a configuration file and it's metadata from the current gocart repo",
	Long: `Delete a configuration file and it's metadata from the current gocart repo
    gocart delete config vimrc`,
	Args: cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}
		gcState, err := gocart.OpenGoCartState()
		if err != nil {
			log.Fatal(err)
		}
		gcState.DeleteConfig(name)
		err = os.Remove(name)
		if err != nil {
			log.Fatal(err)
		}
		err = gcState.Serialize()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("removed config %s, please commit to your VCS\n", name)
		return
	},
}

func init() {
	deleteCmd.AddCommand(configDelCmd)

	var Name string
	configDelCmd.Flags().StringVarP(&Name, "name", "n", "", "config file name")
	configDelCmd.MarkFlagRequired("name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configDelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configDelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
