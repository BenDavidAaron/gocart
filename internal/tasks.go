package gocart

import (
	"errors"
)

type ConfigSpec struct {
	Name     string
	Path     string
	Platform string
}

func InitRepo() error {
	_, err := InitGoCartState()
	return err
}

func InstallRepo() error {
	gcState, err := OpenGoCartState()
	if err != nil {
		return err
	}
	if gcState.Configs == nil {
		return errors.New("gocart: no configs to install")
	}
	for _, cfg := range gcState.Configs {
		err = InsertConfig(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetPlatform() (string, error) {
	var platform string
	gcState, err := OpenGoCartState()
	if err != nil {
		return platform, err
	}
	return gcState.Platform, nil

}

func SetPlatform(newPlatform string) error {
	gcState, err := OpenGoCartState()
	if err != nil {
		return err
	}
	gcState.Platform = newPlatform
	err = gcState.Serialize()
	if err != nil {
		return err
	}
	return nil
}

func GetConfigSpec(cfgName string) (ConfigSpec, error) {
	// Get a config spec from the current dir's gocart mapping
	var cfg ConfigSpec
	gcState, err := OpenGoCartState()
	if err != nil {
		return cfg, err
	}
	cfg = gcState.Configs[cfgName]
	return cfg, nil
}

func GetAllConfigs() ([]ConfigSpec, error) {
	// Get every config spec in the current dir's gocart mapping
	var cfgs []ConfigSpec
	gcState, err := OpenGoCartState()
	if err != nil {
		return cfgs, err
	}
	cfgs = make([]ConfigSpec, 0, len(gcState.Configs))
	for _, cfg := range gcState.Configs {
		cfgs = append(cfgs, cfg)
	}
	return cfgs, nil
}

func PutConfigSpec(cfg ConfigSpec) error {
	// Put a new config spec in the gocart mapping file
	gcState, err := OpenGoCartState()
	if err != nil {
		return err
	}
	err = LinkConfig(cfg)
	if err != nil {
		return err
	}
	gcState.Configs[cfg.Name] = cfg
	err = gcState.Serialize()
	if err != nil {
		return err
	}
	return nil
}

func DeleteConfigSpec(cfgName string) error {
	// Remove a config spec from the current dir's gocart mapping
	gcState, err := OpenGoCartState()
	if err != nil {
		return err
	}
	cfg := gcState.Configs[cfgName]
	err = UnlinkConfig(cfg)
	if err != nil {
		return err
	}
	delete(gcState.Configs, cfgName)
	err = gcState.Serialize()
	if err != nil {
		return err
	}
	return nil
}
