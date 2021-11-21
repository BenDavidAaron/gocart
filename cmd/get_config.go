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
	"strings"

	gocart "github.com/BenDavidAaron/gocart/internal"
	"github.com/spf13/cobra"
)

// configGetCmd represents the configGet command
var configGetCmd = &cobra.Command{
	Use:   "config",
	Short: "Get a config by name, or get all configs",
	Long: `Get a config by name, skipping the name will Get all configs
	cobra configGet vimrc //gets the vimrc mapping
	cobra configGet       //gets all stored configs`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		cfgName := strings.Join(args, "")
		if cfgName == "" {
			var cfgs []gocart.ConfigSpec
			cfgs, err = gocart.GetAllConfigs()
			if err != nil {
				log.Panicf("gocart: could not fetch configs from file", err)
			}
			fmt.Println(cfgs)
		} else {
			var cfg gocart.ConfigSpec
			cfg, err = gocart.GetConfigSpec(cfgName)
			if err != nil {
				log.Panicf("gocart: could not fetch config", cfgName, "from file", err)
			}
			fmt.Println(cfg)
		}
	},
}

func init() {
	getCmd.AddCommand(configGetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configGetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configGetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
