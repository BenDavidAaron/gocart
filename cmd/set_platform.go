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

// platformSetCmd represents the platformSet command
var setPlatformCmd = &cobra.Command{
	Use:   "platform",
	Short: "Set the currently active platform",
	Long: `Set the currently active platform
	gocart platformSet openbsd  // Sets the platform to openbsd`,
	Run: func(cmd *cobra.Command, args []string) {
		err := gocart.SetPlatform(args[0])
		if err != nil {
			log.Fatalf("gocart: could not write platform to disk", err)
		}
		fmt.Printf("Set Platform to:", args[0])
	},
}

func init() {
	setCmd.AddCommand(setPlatformCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// platformSetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// platformSetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
