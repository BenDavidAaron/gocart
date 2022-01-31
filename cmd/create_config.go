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

// configPutCmd represents the configPut command
var configPutCmd = &cobra.Command{
	Use:   "config",
	Short: "Create a new config file in the gocart repo and track it",
	Long: `Create a new config gile in the gocart repo and track it.

	this will create a file with the specified name name specified in the gocart 
	repo, from here it can be edited, and linked to the appropriate location on 
	your system.

    gocart create config vimrc`,
	Args: cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configPut called")
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}

		gcState, err := gocart.OpenGoCartState()
		if err != nil {
			log.Panicf("gocart: unable to load application data from disk", err)
		}
		var cfg gocart.ConfigSpec
		if cfg, ok := gcState.Configs[name]; ok {
			fmt.Printf("%s is already present in this repo, exiting")
			return
		} else {
			cfg.Init()
			cfg.Name = name
			os.Touch(name)
		}
		gcState.PutConfig(cfg)
		err = gcState.Serialize()
		if err != nil {
			log.Panicf("gocart: unable to write config to disk", err)
		}
		fmt.Printf("Saved %s to gocart\n", cfg.Name)
	},
}

func init() {
	addCmd.AddCommand(configPutCmd)

	var Name string
	configPutCmd.Flags().StringVarP(&Name, "name", "n", "", "config file name")
	configPutCmd.MarkFlagRequired("name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configPutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configPutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
