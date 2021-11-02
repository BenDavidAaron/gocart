package gocart

type ConfigSpec struct {
	Name     string
	Path     string
	Platform string
}

func InitRepo() error {
	newStore := new(GoCartStore)
	newStore.Path = MappingFilePath
	err := newStore.Init()
	return err
}

func GetPlatform() (string, error) {
	var platform string
	gcState, err := ReadGoCartState()
	if err != nil {
		return platform, err
	}
	return gcState.Platform, nil

}

func SetPlatform(newPlatform string) error {
	gcState, err := ReadGoCartState()
	if err != nil {
		return err
	}
	gcState.Platform = newPlatform
	err = WriteGocartState(gcState)
	if err != nil {
		return err
	}
	return nil
}

func GetConfigSpec(cfgName string) (ConfigSpec, error) {
	// Get a config spec from the current dir's gocart mapping
	var cfg ConfigSpec
	gcState, err := ReadGoCartState()
	if err != nil {
		return cfg, err
	}
	cfg = gcState.configs[cfgName]
	return cfg, nil
}

func PutConfigSpec(cfg ConfigSpec) error {
	// Put a new config spec in the gocart mapping file
	gcState, err := ReadGoCartState()
	if err != nil {
		return err
	}
	// TODO: Copy config file to workdir
	// TODO: Backlink config file from workdir to cfg.Path
	gcState.configs[cfg.Name] = cfg
	err = WriteGocartState(gcState)
	if err != nil {
		return err
	}
	return nil
}

func DeleteConfigSpec(cfgName string) error {
	// Remove a config spec from the current dir's gocart mapping
	gcState, err := ReadGoCartState()
	if err != nil {
		return err
	}
	// TODO: move config file to cfg.Path
	// TODO: remove backlink from repo's config file
	delete(gcState.configs, cfgName)
	err = WriteGocartState(gcState)
	if err != nil {
		return err
	}
	return nil
}
