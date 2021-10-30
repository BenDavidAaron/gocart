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

	"github.com/spf13/cobra"
)

// configPutCmd represents the configPut command
var configPutCmd = &cobra.Command{
	Use:   "configPut",
	Short: "Put a Config File into the current gocart mapping",
	Long: `Put a Config File into the current gocart mapping

	this will create a file at the name specified in the gocart repo, 
	copy the config file at the supplied path into the new file,
	and then replace the old config file with a symlink to the new config file in the gocart repo

    gocart configPut vimrc ~/.vimrc`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configPut called")
	},
}

func init() {
	rootCmd.AddCommand(configPutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configPutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configPutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
