package gocart

import (
	"testing"
)

func TestSetAndGetPlatform(t *testing.T) {
	firstPlatform, err := GetPlatform()
	if err != nil {
		t.Error(err)
	}
	err = SetPlatform("TestOS")
	if err != nil {
		t.Error(err)
	}
	secondPlatform, err := GetPlatform()
	if err != nil {
		t.Error(err)
	}
	if firstPlatform == secondPlatform {
		t.Errorf("The Platform Setting was not changed on update, %s - %s", firstPlatform, secondPlatform)
	}
}
