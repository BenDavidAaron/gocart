package gocart

type ConfigSpec struct {
	name     string
	path     string
	platform string
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
