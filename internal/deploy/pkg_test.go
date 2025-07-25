package deploy

import (
	"errors"
	errors2 "hephaestus/internal/errors"
	"testing"
)

func TestLoadPkg(t *testing.T) {
	pkg, err := LoadPkg("../../test/test_invalid.zip")
	if err != nil {
		if errors.Is(err, errors2.ErrFileNoEqual) {
			t.Log("test invalid success")
		} else {
			t.Fatal(err)
		}
	}
	t.Logf("%v\n", pkg)
}

func TestLoadPkgValid(t *testing.T) {
	pkg, err := LoadPkg("../../test/test_valid.zip")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("test valid success")
	t.Logf("%v\n", pkg)
}

func TestUnzipPkg(t *testing.T) {
	err := Load("../../test/test.zip", "test")
	if err != nil {
		t.Fatal(err)
	}
}
