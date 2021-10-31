package gocart

import (
	"encoding/json"
	"io/ioutil"
)

const MappingFilePath string = ".gocart.json"

// Go Cart Data Store
type GoCartStore struct {
	Path string
}

func (gcStore *GoCartStore) Init() error {
	// Create a blank GoCart state and serialize it to disk
	err := gcStore.Serialize(GoCartState{map[string]ConfigSpec{}, ""})
	if err != nil {
		return err
	}
	return nil
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
func InitGoCartState() (GoCartState, error) {
	//Create  a new gocart repo in the current directory and return the state as a mutable object
	store := GoCartStore{
		Path: MappingFilePath,
	}
	state, err := store.Deserialize()
	if err != nil {
		return GoCartState{}, err
	}
	return state, err
}

func ReadGoCartState() (GoCartState, error) {
	//Open a the local repo's mapping and return the state as a mutable object
	store := GoCartStore{
		Path: MappingFilePath,
	}
	state, err := store.Deserialize()
	if err != nil {
		return GoCartState{}, err
	}
	return state, err
}

func WriteGocartState(gcState GoCartState) error {
	//Write modified state to the local repo's mapping
	store := GoCartStore{Path: MappingFilePath}
	err := store.Serialize(gcState)
	if err != nil {
		return err
	}
	return nil
}

type GoCartState struct {
	configs  map[string]ConfigSpec
	Platform string
}

func (gcState *GoCartState) PutConfig(cfg ConfigSpec) {
	gcState.configs[cfg.name] = cfg
	return
}

func (gcState *GoCartState) GetConfig(name string) ConfigSpec {
	return gcState.configs[name]
}

func (gcState *GoCartState) GetConfigs() map[string]ConfigSpec {
	return gcState.configs
}

func (gcState *GoCartState) GetPlatform() string {
	return gcState.Platform
}

func (gcState *GoCartState) SetPlatform(newPlatform string) {
	gcState.Platform = newPlatform
}
