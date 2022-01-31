package gocart

import (
	"errors"
	"os"
	"path/filepath"
)

func link(srcName, targetPath string) error {
	// Take a config file name (from working dir) and a target path,
	// and insert a symlink to ./name located at path
	gcCfgPath, err := filepath.Abs(srcName)
	if err != nil {
		return err
	}
	if !filepath.IsAbs(targetPath) {
		return errors.New("gocart: cannot Link non-absolute path")
	}
	err = os.Symlink(gcCfgPath, targetPath)
	if err != nil {
		return err
	}
	return nil
}

func unlink(name, targetPath string) error {
	// Delete the symlink at targetPath,
	// copy ./name to path
	ourPath, err := filepath.Abs(targetPath)
	if err != nil {
		return err
	}
	if !filepath.IsAbs(Path) {
		return errors.New("gocart: cannot Unlink to a non-absolute path")
	}
	err = os.Remove(Path) // Delete the symlink at Path
	if err != nil {
		return err
	}
	return nil
}
