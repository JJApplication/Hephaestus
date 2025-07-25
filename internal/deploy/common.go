package deploy

import (
	copy2 "github.com/otiai10/copy"
	"gopkg.in/yaml.v3"
	"hephaestus/internal/config"
	"hephaestus/internal/constant"
	"hephaestus/internal/errors"
	"os"
	"path/filepath"
	"time"
)

// 部署相关的目录 TMP APP BACK
// TMP上传部署包的目录，按照服务名存储$TMP/$APP
// APP部署包解压的目录，新服务的数据目录 $APP/$APP 同一时刻只存在一个服务包 所以解压前做删除处理
// BACK微服务备份目录 $BACK/$APP 备份目录名称使用年月日时分秒创建 最多只允许存在5个备份目录

// GetDeployTmpPath 获取部署解压临时目录
func GetDeployTmpPath() string {
	return filepath.Join(config.LoadConfig().DeployPath, constant.TempUpload)
}

// GetDeployDataPath 获取部署数据目录
// 所有版本数据都存储在此目录下
func GetDeployDataPath() string {
	return filepath.Join(config.LoadConfig().DeployPath, constant.AppDeploy)
}

// GetAppPath 获取微服务运行目录
func GetAppPath(app string) (string, error) {
	if app == "" {
		return "", errors.ErrEmptyApp
	}
	return filepath.Join(config.LoadConfig().AppRoot, app), nil
}

// GetDeployAppPath 获取部署数据目录下的微服务数据目录 即解压后的微服务数据目录
func GetDeployAppPath(app string) (string, error) {
	if config.LoadConfig().DeployPath == "" {
		return "", errors.ErrEmptyDeploy
	}
	if app == "" {
		return "", errors.ErrEmptyApp
	}
	return filepath.Join(GetDeployDataPath(), app), nil
}

// GetDeployBackupPath 获取微服务的备份目录
func GetDeployBackupPath(app string) (string, error) {
	if config.LoadConfig().DeployPath == "" {
		return "", errors.ErrEmptyDeploy
	}
	if app == "" {
		return "", errors.ErrEmptyApp
	}
	return filepath.Join(config.LoadConfig().DeployPath, constant.BackUp, app), nil
}

// UnzipDeployPkg 解压部署包 创建数据目录
func UnzipDeployPkg(app string) error {
	tmpUpload := GetDeployTmpPath()
	deployPath := GetDeployDataPath()
	if deployPath == "" {
		return errors.ErrEmptyDeploy
	}
	if app == "" {
		return errors.ErrEmptyApp
	}
	// 解压微服务下的pkg文件到指定微服务
	appPkgPath := filepath.Join(tmpUpload, app, constant.PkgFile)
	if _, err := os.Stat(appPkgPath); os.IsNotExist(err) {
		return err
	}
	appDeployPath := filepath.Join(deployPath, app)
	if _, err := os.Stat(appDeployPath); os.IsNotExist(err) {
		_ = os.MkdirAll(appDeployPath, os.ModePerm)
	} else {
		_ = os.RemoveAll(appDeployPath)
		_ = os.MkdirAll(appDeployPath, os.ModePerm)
	}
	return Load(appPkgPath, appDeployPath)
}

func GenerateRollback(app string, pkg *Pkg) error {
	backupDate := time.Now().Format("20060102150405")
	appBackPath, err := GetDeployBackupPath(app)
	if err != nil {
		return err
	}
	if _, err = os.Stat(appBackPath); os.IsNotExist(err) {
		_ = os.MkdirAll(appBackPath, os.ModePerm)
	}
	// 存在才存储
	if _, err = os.Stat(appBackPath); err != nil {
		return err
	}

	// 创建日期备份
	backDir := filepath.Join(appBackPath, backupDate)
	_ = os.MkdirAll(backDir, os.ModePerm)

	// 先创建回滚文件
	var rollback Rollback
	for file, pkgType := range pkg.Pkg {
		switch pkgType {
		case "add":
			rollback.Add = append(rollback.Add, file)
		case "replace":
			rollback.Modify = append(rollback.Modify, file)
		case "delete":
			rollback.Delete = append(rollback.Delete, file)
		}
	}

	rollbackFile := filepath.Join(backDir, constant.ROLLBACK_FILE)
	data, err := yaml.Marshal(&rollback)
	if err != nil {
		return err
	}
	if err = os.WriteFile(rollbackFile, data, os.ModePerm); err != nil {
		return err
	}

	// 开始基于pkg文件进行备份任务
	appPath, err := GetAppPath(app)
	if err != nil {
		return err
	}
	for file, pkgType := range pkg.Pkg {
		if pkgType == "replace" || pkgType == "delete" {
			// 复制文件到last
			if err = copy2.Copy(filepath.Join(appPath, file), filepath.Join(backDir, file)); err != nil {
				return err
			}
		}
	}

	return nil
}
