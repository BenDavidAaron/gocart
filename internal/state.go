package gocart

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

const MappingFilePath string = "./.gocart.json"

//Go Cart Application State
type GoCartState struct {
	Configs  map[string]ConfigSpec
	Platform string
	Path     string
}

func InitGoCartState() (GoCartState, error) {
	//Create a new gocart repo in the current directory and initialize it with empty state
	gcState := GoCartState{
		Path:     "",
		Configs:  make(map[string]ConfigSpec),
		Platform: "",
	}
	var err error
	gcState.Path, err = filepath.Abs(MappingFilePath)
	err = gcState.Serialize()
	return gcState, err
}

func OpenGoCartState() (GoCartState, error) {
	// Load the gocart repo in the current directory and return it
	var gcState GoCartState
	var err error
	gcState.Path, err = filepath.Abs(MappingFilePath)
	if err != nil {
		return gcState, err
	}
	gcState, err = gcState.Deserialize()
	return gcState, err
}

func (gcState *GoCartState) Serialize() error {
	// Read the file at ./.gocart.json and load it's platform name and configs into memory
	stateData, err := json.Marshal(gcState)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(gcState.Path, stateData, 0640)
	if err != nil {
		return err
	}
	return nil
}

func (gcStore *GoCartState) Deserialize() (GoCartState, error) {
	// Save the configs and platform name in memory to a file at ./.gocart.json
	var state GoCartState
	kvFile, err := ioutil.ReadFile(MappingFilePath)
	if err != nil {
		return state, err
	}
	err = json.Unmarshal([]byte(kvFile), &state)
	if err != nil {
		return state, err
	}
	if state.Configs == nil {
		state.Configs = make(map[string]ConfigSpec)
	}
	return state, nil
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

func (gcState *GoCartState) GetPlatform() string {
	return gcState.Platform
}

func (gcState *GoCartState) SetPlatform(newPlatform string) {
	gcState.Platform = newPlatform
}

//Go Cart Configuration Specification
type ConfigSpec struct {
	Name  string
	Paths map[string]string
}

func MakeConfigSpec() ConfigSpec {
	cfg := new(ConfigSpec)
	cfg.Paths = map[string]string{}
	return *cfg
}

func (cfg *ConfigSpec) AddPath(platform, path string) {
	// Add a Platform Specific path to this config
	cfg.Paths[platform] = path
}

func (cfg *ConfigSpec) RemovePath(platform string) {
	// Remove a Platform Specific path from this config
	delete(cfg.Paths, platform)
}
