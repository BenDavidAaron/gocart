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
	"strings"

	gocart "github.com/BenDavidAaron/gocart/internal"
	"github.com/spf13/cobra"
)

// configDelCmd represents the configDel command
var deleteCfgCmd = &cobra.Command{
	Use:   "config",
	Short: "Delete a Configuration from the current gocart repo",
	Long: `Delete a Congig file from the current gocart repo and restore the config file to it's original home
    gocart configDel vimrc`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configDel called")
		var name string
		var err error
		fmt.Println(args)
		name = strings.Join(args, "")
		fmt.Println(name)
		err = gocart.DeleteConfigSpec(name)
		if err != nil {
			fmt.Println(err)
		}
		return
	},
}

func init() {
	deleteCmd.AddCommand(deleteCfgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configDelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configDelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
