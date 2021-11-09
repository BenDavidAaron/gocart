package gocart

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LinkConfig(cfg ConfigSpec) error {
	// Take a config Spec, copy it into the gocart repo,
	// and create a backlink at the original location
	err := link(cfg.Name, cfg.Path)
	if err != nil {
		return err
	}
	return nil
}
func link(name, path string) error {
	// Take a config file name and a path,
	//   copy the file at path to ./name,
	//   and insert a symlink to ./name located at path
	newPath, err := filepath.Abs(name)
	if err != nil {
		return err
	}
	if !filepath.IsAbs(path) {
		return errors.New("gocart: cannot Link non-absolute path")
	}
	err = copyFile(path, newPath)
	if err != nil {
		return err
	}
	err = os.Remove(path)
	if err != nil {
		return err
	}
	err = os.Symlink(newPath, path)
	if err != nil {
		return err
	}
	return nil
}

func UnlinkConfig(cfg ConfigSpec) error {
	// Take a config Spec, Copy it to it's original location,
	// overwriting the backlink and remove the copy of the cfg from the working dir
	err := unlink(cfg.Name, cfg.Path)
	if err != nil {
		return err
	}
	return nil
}
func unlink(name, Path string) error {
	// Delete the symlink at path,
	// copy ./name to path
	ourPath, err := filepath.Abs(name)
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
	err = copyFile(ourPath, Path)
	if err != nil {
		return err
	}
	err = os.Remove(ourPath)
	if err != nil {
		return err
	}
	return nil
}
func copyFile(src string, dst string) error {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func InsertConfig(cfg ConfigSpec) error {
	// Take a config spec, and insert a symlink to the repo's copy
	// of that config file at the config spec's path.
	repoCfgPath, err := filepath.Abs(cfg.Name)
	if err != nil {
		return err
	}
	err = insert(repoCfgPath, cfg.Path)
	return err
}

func insert(repoPath, systemPath string) error {
	// Place a symlink pointing from path to ./name
	err := os.Symlink(repoPath, systemPath)
	if err != nil {
		return err
	}
	return nil
}
