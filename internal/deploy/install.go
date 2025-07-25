package deploy

import (
	copy2 "github.com/otiai10/copy"
	"os"
	"path/filepath"
)

// 安装微服务

func Install(pkg Pkg, app string) error {
	newAppDataPath, err := GetDeployAppPath(app)
	if err != nil {
		return err
	}
	// 解压服务包
	if err = UnzipDeployPkg(app); err != nil {
		return err
	}

	appPath, err := GetAppPath(app)
	if err != nil {
		return err
	}
	// 进行数据备份替换 pkg中定义的文件路径都是基于微服务根目录的
	// 先进行全量备份
	if err = GenerateRollback(app, &pkg); err != nil {
		return err
	}

	// 进行文件替换
	for file, pkgType := range pkg.Pkg {
		switch pkgType {
		case "replace":
			if err := os.Remove(filepath.Join(appPath, file)); err != nil {
				return err
			}
			if err := copy2.Copy(filepath.Join(newAppDataPath, file), filepath.Join(appPath, file)); err != nil {
				return err
			}
		case "delete":
			if err := os.Remove(filepath.Join(appPath, file)); err != nil {
				return err
			}
		case "add":
			_, err = os.Stat(filepath.Join(appPath, file))
			if err == nil {
				continue
			} else {
				if os.IsNotExist(err) {
					if err := copy2.Copy(filepath.Join(newAppDataPath, file), filepath.Join(appPath, file)); err != nil {
						return err
					}
				}
			}
		default:

		}
	}

	return nil
}
