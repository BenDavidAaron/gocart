package gocart

type ConfigSpec struct {
	name     string
	path     string
	platform string
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
