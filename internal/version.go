package gocart

import (
	"fmt"
	"strconv"
)

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

func ParseVersionString(v_str string) (gocartVersion, error) {
	v_chars := strings.split(v_str, ".")
	v_nums := make([]int, len(v_chars))
	for i, s := range v_chars {
		v_nums[i], _ = strconv.Atoi(s)
	}
	vers := gocartVersion{}
	vers.Major = v_nums[0]
	vers.Minor = v_nums[1]
	vers.Patch = v_nums[2]
	return vers
}
