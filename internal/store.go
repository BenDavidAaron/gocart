package gocart

import (
	"encoding/json"
	"io/ioutil"
)

const MappingFilePath string = ".gocart.json"

type ConfigSpec struct {
	name     string
	path     string
	platform string
}

// Go Cart Data Store
type GoCartStore struct {
	Path string
}

func (gcStore *GoCartStore) Serialize(gcState GoCartState) error {
	stateData, err := json.Marshal(gcState)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(gcStore.Path, stateData, 0640)
	if err != nil {
		return err
	}
	return nil
}

func (gcStore *GoCartStore) Deserialize() (GoCartState, error) {
	var state GoCartState
	kvFile, err := ioutil.ReadFile(MappingFilePath)
	if err != nil {
		return state, err
	}
	err = json.Unmarshal([]byte(kvFile), &state)
	if err != nil {
		return state, err
	}
	return state, nil
}

//Go Cart Application State
type GoCartState struct {
	configs  map[string]ConfigSpec
	Platform string
}

func (gcStore *GoCartState) Put(cfg ConfigSpec) {
	gcStore.configs[cfg.name] = cfg
	return
}

func (gcStore *GoCartState) Get(name string) ConfigSpec {
	return gcStore.configs[name]
}

func (gcStore *GoCartState) GetAll() map[string]ConfigSpec {
	return gcStore.configs
}
