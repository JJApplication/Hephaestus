package deploy

import (
	copy2 "github.com/otiai10/copy"
	"gopkg.in/yaml.v3"
	"hephaestus/internal/constant"
	"os"
	"path/filepath"
)

// RollbackInstall 回滚
func RollbackInstall(app string, version string) error {
	// 回滚时已保证微服务处于停止态
	// 解析回滚配置
	backupPath, err := GetDeployBackupPath(app)
	if err != nil {
		return err
	}
	backupVersion := filepath.Join(backupPath, version)
	if _, err = os.Stat(backupVersion); os.IsNotExist(err) {
		return err
	}
	var rb Rollback
	rbFile := filepath.Join(backupPath, version, constant.ROLLBACK_FILE)
	if _, err = os.Stat(rbFile); os.IsNotExist(err) {
		return err
	}
	data, err := os.ReadFile(rbFile)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(data, &rb); err != nil {
		return err
	}

	// 根据配置处理文件
	appPath, err := GetAppPath(app)
	if err != nil {
		return err
	}
	// 由于文件处理导致的失败 直到最后再处理
	// 如果最终有错误 实际文件已经被处理，微服务回滚视作完成
	for _, s := range rb.Add {
		// 新增文件要删除（v1版本不处理多版本导致的差异，如果旧版本新增文件 在新版本实际属于修改 那么仍视作新增）
		if s != "" {
			_ = os.RemoveAll(filepath.Join(appPath, s))
		}
	}

	for _, s := range rb.Modify {
		sPath := filepath.Join(backupVersion, s)
		if _, err = os.Stat(sPath); os.IsNotExist(err) {
			continue
		}
		// 存在则替换
		_ = os.RemoveAll(filepath.Join(appPath, s))
		_ = copy2.Copy(sPath, filepath.Join(appPath, s))
	}

	for _, s := range rb.Delete {
		sPath := filepath.Join(backupVersion, s)
		if _, err = os.Stat(sPath); os.IsNotExist(err) {
			continue
		}
		// 存在则替换
		_ = os.RemoveAll(filepath.Join(appPath, s))
		_ = copy2.Copy(sPath, filepath.Join(appPath, s))
	}

	return nil
}

// RollbackAfter 清理工作
// 回滚后 原有的旧版服务目录已经不再需要
func RollbackAfter(app string, version string) error {
	backupPath, err := GetDeployBackupPath(app)
	if err != nil {
		return err
	}
	if _, err = os.Stat(backupPath); os.IsNotExist(err) {
		return nil
	}
	backupVersion := filepath.Join(backupPath, version)

	return os.RemoveAll(backupVersion)
}
