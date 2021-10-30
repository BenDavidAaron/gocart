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
type GoCartState struct {
	configs  map[string]ConfigSpec
	Platform string
}

type KeyValueStore struct {
	Path string
}

func (*KeyValueStore) Serialize(state GoCartState) error {
	stateData, err := json.Marshal(state)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(MappingFilePath, stateData, 0640)
	if err != nil {
		return err
	}
	return nil
}

func (*KeyValueStore) Deserialize() (GoCartState, error) {
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

func (kvStore *KeyValueStore) Put(cfg ConfigSpec) error {
	kvState, err := kvStore.Deserialize()
	if err != nil {
		return err
	}
	kvState.configs[cfg.name] = cfg
	err = kvStore.Serialize(kvState)
	if err != nil {
		return err
	}
	return nil
}

func (kvStore *KeyValueStore) Get(name string) (ConfigSpec, error) {
	kvState, err := kvStore.Deserialize()
	if err != nil {
		return ConfigSpec{}, err
	}
	return kvState.configs[name], nil
}

func (kvStore *KeyValueStore) GetAll() (map[string]ConfigSpec, error) {
	cfgs := make(map[string]ConfigSpec)
	kvState, err := kvStore.Deserialize()
	if err != nil {
		return cfgs, err
	}
	return kvState.configs, nil
}
