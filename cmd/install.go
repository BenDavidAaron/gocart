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
	"log"

	gocart "github.com/BenDavidAaron/gocart/internal"
	"github.com/spf13/cobra"
)

// repoInstallCmd represents the repoInstall command
var repoInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the current gocart repo's config on your system",
	Long: `Insert a symlink for each config in the current gocart repo
	pointing from each config's ./name to it's path on your system`,
	Run: func(cmd *cobra.Command, args []string) {
		err := gocart.InstallRepo()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Installed configs from workdir and .gocart.json")
	},
}

func init() {
	rootCmd.AddCommand(repoInstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoInstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoInstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
