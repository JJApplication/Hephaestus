package deploy

import (
	"hephaestus/internal/constant"
	"path/filepath"
	"testing"
)

const (
	APP  = "test"
	APP2 = "test_deep"
)

func TestInstall(t *testing.T) {
	t.Log("test install", APP)
	pkgPath := GetDeployTmpPath()
	t.Log("pkgPath:", filepath.Join(pkgPath, APP))
	appPkgData, err := LoadPkg(filepath.Join(pkgPath, APP, constant.PkgFile))
	if err != nil {
		t.Error(err)
		return
	}
	err = Install(*appPkgData, APP)
	if err != nil {
		t.Error(err)
	}
	t.Log("test install success")
}

func TestInstallDeep(t *testing.T) {
	t.Log("test install", APP2)
	pkgPath := GetDeployTmpPath()
	t.Log("pkgPath:", filepath.Join(pkgPath, APP2))
	appPkgData, err := LoadPkg(filepath.Join(pkgPath, APP2, constant.PkgFile))
	if err != nil {
		t.Error(err)
		return
	}
	err = Install(*appPkgData, APP2)
	if err != nil {
		t.Error(err)
	}
	t.Log("test install success")
}
