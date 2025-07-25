package deploy

import (
	copy2 "github.com/otiai10/copy"
	"testing"
)

// 测试复制行为

func TestCpFile2File(t *testing.T) {
	copy2.Copy("../../test/cp/file1", "../../test/cp/file2")
}

func TestCpFile2Dir(t *testing.T) {
	copy2.Copy("../../test/cp/file1", "../../test/cp/dir1/file1")
}

func TestCpDir2Dir(t *testing.T) {
	copy2.Copy("../../test/cp/dir1", "../../test/cp/dir2")
}

func TestCpDir2ParentDir(t *testing.T) {
	// 存在则失败
	copy2.Copy("../../test/cp/dir1", "../../test/cp")
}
