package gocart

import (
	"path/filepath"
	"testing"
)

var testCfg ConfigSpec

func TestLinkConfig(t *testing.T) {
	testCfg = ConfigSpec{
		Name:     filepath.Join(t.TempDir(), "testTarget"),
		Path:     filepath.Join(t.TempDir(), "testrc"),
		Platform: "test",
	}
	err := LinkConfig(testCfg)
	if err != nil {
		t.Error(err)
	}
	resolved_symlink, err := filepath.Abs(testCfg.Path)
	testPath, err := filepath.Abs(testCfg.Name)
	if resolved_symlink != testPath {
		t.Errorf("Testing Linker failed, expected %s got %s", testPath, resolved_symlink)
	}
}
