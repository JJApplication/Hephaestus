package deploy

import (
	"gopkg.in/yaml.v3"
)

// Rollback 记录新增，删除，修改项
// 新增的删除
// 删除的恢复
// 更新的替换
// 文件名发生变化的视作新增或删除
type Rollback struct {
	Add    []string `yaml:"add"`
	Delete []string `yaml:"delete"`
	Modify []string `yaml:"modify"`
}

func (r *Rollback) String() []byte {
	d, err := yaml.Marshal(r)
	if err != nil {
		return nil
	}
	return d
}
