package gocart

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

func (cfg *ConfigSpec) Link(platformName string) error {
	// Link a config file into the system using the specified platform
	err := link(cfg.Name, cfg.Paths[platformName])
	if err != nil {
		return err
	}
	return nil
}

func (cfg *ConfigSpec) Unlink(platformName string) error {
	// Unlink a config file from the system using the specified platform
	err := unlink(cfg.Name, cfg.Paths[platformName])
	if err != nil {
		return err
	}
	return nil
}
