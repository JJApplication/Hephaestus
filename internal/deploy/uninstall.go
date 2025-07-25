package deploy

import "os"

// Uninstall 卸载 不做任何删除操作主要依赖Apollo
func Uninstall(app string) {

}

func ForceUninstall(app string) error {
	// 保证服务已经停止
	// 删除所有相关的资源
	var err error
	appPath, err := GetAppPath(app)
	if err == nil {
		_ = os.RemoveAll(appPath)
	}
	deployAppPath, err := GetDeployAppPath(app)
	if err == nil {
		_ = os.RemoveAll(deployAppPath)
	}
	backupAppPath, err := GetDeployBackupPath(app)
	if err == nil {
		_ = os.RemoveAll(backupAppPath)
	}

	return err
}
