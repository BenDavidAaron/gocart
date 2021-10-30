package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	"github.com/google/subcommands"
)

type ConfigSpec struct {
	name     string
	path     string
	platform string
}

type GoCartState struct {
	configs  map[string]ConfigSpec
	platform string
}

const mappingFilePath string = ".gocart.map"

type KeyValueStore struct {
	path string
}

func (*KeyValueStore) Serialize(state GoCartState) error {
	stateData, err := json.Marshal(state)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(mappingFilePath, stateData, 0640)
	if err != nil {
		return err
	}
	return nil
}

func (*KeyValueStore) Deserialize() (GoCartState, error) {
	var state GoCartState
	kvFile, err := ioutil.ReadFile(mappingFilePath)
	if err != nil {
		return state, err
	}
	err = json.Unmarshal([]byte(kvFile), &state)
	if err != nil {
		return state, err
	}
	return state, nil
}

func (kvStore *KeyValueStore) Put(cfg ConfigSpec) error {
	kvState, err := kvStore.Deserialize()
	if err != nil {
		return err
	}
	kvState.configs[cfg.name] = cfg
	err = kvStore.Serialize()
	if err != None {
		return err
	}
	return nil
}

type initCmd struct {
}

func (*initCmd) Name() string     { return "init" }
func (*initCmd) Synopsis() string { return "create a gocart mapping file in the current directory" }
func (*initCmd) Usage() string    { return "init" }

func (cmd *initCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	os.Create(mappingFilePath)
	return subcommands.ExitSuccess
}
func (cmd *initCmd) SetFlags(f *flag.FlagSet) {
}

type addCmd struct {
	config_name      string
	config_file_path string
}

func (*addCmd) Name() string     { return "add" }
func (*addCmd) Synopsis() string { return "add a config file to the mapping in the current directory" }
func (*addCmd) Usage() string {
	return `add <config_file_path> <config_name>
    add a config file to the mapping in the current directory
	`
}

func (cmd *addCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// open config file and insert a new key:value pair
	return subcommands.ExitSuccess
}
func (cmd *addCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.config_file_path, "p", "", "path to config file to track")
	f.StringVar(&cmd.config_name, "n", "", "name of a config")
}

type removeCmd struct {
	config_name string
}

func (*removeCmd) Name() string { return "remove" }
func (*removeCmd) Synopsis() string {
	return "remove a config file from the local mapping and restore it to it's directory"
}
func (*removeCmd) Usage() string {
	return `remove <config_name>
	remove a config file from the local mapping and restore it to it's directory
	`
}

func (cmd *removeCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// open config file and insert a new key:value pair
	return subcommands.ExitSuccess
}
func (cmd *removeCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.config_name, "config_name", "", "name of the config file to untrack")
}

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&initCmd{}, "")
	subcommands.Register(&addCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
