package gocart

import "fmt"

const (
	versionMajor int = 0
	versionMinor int = 0
	versionPatch int = 0
)

type gocartVersion struct {
	Major int
	Minor int
	Patch int
}

func GetVersion() *gocartVersion {
	v := gocartVersion{}
	v.Major = versionMajor
	v.Minor = versionMinor
	v.Patch = versionPatch
	return &v
}

func GetVersionString() string {
	v := GetVersion()
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
