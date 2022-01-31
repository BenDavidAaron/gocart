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

// repoInitCmd represents the repoInit command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a blank mapping file in the current directory",
	Long: `Create a blank mapping file in the current directory. Run this
	in an empty git repository so you can check in config files scattered 
	around your filesystem`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := gocart.LoadGoCartState()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Created new gocart repo in working dir")
		fmt.Println("Please add '.gocart.json' to version control")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoInitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoInitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
