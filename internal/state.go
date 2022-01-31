package gocart

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const MappingFilePath string = "./gocart.json"

//Go Cart Application State
type GoCartState struct {
	Configs        map[string]ConfigSpec
	ActivePlatform string
	Platforms      []string
	Version        string
}

func NewGoCartState() *GoCartState {
	//Create a new gocart repo in the current directory and initialize it with empty state
	gcState := new(GoCartState)
	gcState.instantiate()
	return gcState
}

func (gcState *GoCartState) instantiate() {
	if gcState.Configs == nil {
		gcState.Configs = make(map[string]ConfigSpec)
	}
	if gcState.Platforms == nil {
		gcState.Platforms = make([]string, 0)
	}
	if gcState.Version == "" {
		gcState.Version = GetVersionString()
	}
}

func LoadGoCartState() (GoCartState, error) {
	// Load the gocart repo in the current directory and return it
	var gcState GoCartState
	gcFile, err := ioutil.ReadFile(MappingFilePath)
	if err != nil {
		return gcState, err
	}
	err = json.Unmarshal([]byte(gcFile), &gcState)
	if err != nil {
		return gcState, err
	}
	gcState.instantiate()
	return gcState, nil
}

func (gcState *GoCartState) Save() error {
	stateData, err := json.Marshal(gcState)
	path, err := filepath.Abs(MappingFilePath)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, stateData, 0640)
	if err != nil {
		return err
	}
	return nil
}

func (gcState *GoCartState) PutConfig(cfg ConfigSpec) {
	gcState.Configs[cfg.Name] = cfg
	return
}

func (gcState *GoCartState) GetConfig(name string) ConfigSpec {
	return gcState.Configs[name]
}

func (gcState *GoCartState) GetConfigs() map[string]ConfigSpec {
	return gcState.Configs
}

func (gcState *GoCartState) DeleteConfig(name string) {
	// Delete a config spec from the gocart repo
	delete(gcState.Configs, name)
}

func (gcState *GoCartState) GetPlatform() string {
	return gcState.Platform
}

func (gcState *GoCartState) SetPlatform(newPlatform string) {
	gcState.Platform = newPlatform
}
