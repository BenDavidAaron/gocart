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

	gocart "github.com/BenDavidAaron/gocart/internal"
	"github.com/spf13/cobra"
)

// platformGetCmd represents the platformGet command
var getPlatformCmd = &cobra.Command{
	Use:   "platform",
	Short: "Get the name of the currently selected platform (BSD, Linux, OSX",
	Long: `Get the name of the currently selected platform (BSD, Linux, OSX)
    cobra platformGet  // bsd`,
	Run: func(cmd *cobra.Command, args []string) {
		platform, err := gocart.GetPlatform()
		if err != nil {
			log.Fatal("gocart: failed to get platform from disk", err)
		}
		fmt.Println(platform)
	},
}

func init() {
	getCmd.AddCommand(getPlatformCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// platformGetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// platformGetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
