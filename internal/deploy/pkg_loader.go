package deploy

import (
	"archive/zip"
	"gopkg.in/yaml.v3"
	"hephaestus/internal/constant"
	"hephaestus/internal/errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 加载pkg包

func LoadPkg(pkg string) (*Pkg, error) {
	file, err := zip.OpenReader(pkg)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	// 防止过载
	if len(file.File) < 1 || len(file.File) > constant.MAX_FILES {
		return nil, errors.ErrFileSize
	}
	var valid bool
	var pkgIndex int
	for index, f := range file.File {
		if f.Name == constant.PKG_FILE {
			pkgIndex = index
			valid = true
		}
	}
	if !valid {
		return nil, errors.ErrFileInvalid
	}

	pkgFile := file.File[pkgIndex]
	data, err := pkgFile.Open()
	if err != nil {
		return nil, err
	}
	defer data.Close()
	content, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}

	var pkgData Pkg
	if err = yaml.Unmarshal(content, &pkgData); err != nil {
		return nil, err
	}

	if err = Precheck(&pkgData, file); err != nil {
		return nil, err
	}
	return &pkgData, nil
}

// Precheck 在加载pkg时调用 校验pkg配置和包内文件对应
func Precheck(pkg *Pkg, f *zip.ReadCloser) error {
	// 仅根据pkg文件判断pkg中没有的 压缩包的对应目录和配置目录应该都基于根目录
	pkgFile := pkg.Pkg

	allChecked := true
	for file, pkgType := range pkgFile {
		if pkgType == "replace" || pkgType == "add" {
			// 标准化路径参数统一使用/
			realPkgName := strings.ReplaceAll(file, "\\", "/")
			var checked bool
			for _, zipFile := range f.File {
				realZipName := strings.ReplaceAll(zipFile.Name, "\\", "/")
				if realPkgName == realZipName {
					checked = true
				}
			}

			if !checked {
				allChecked = false
			}
		}
	}

	if !allChecked {
		return errors.ErrFileNoEqual
	}

	return nil
}

// Load 正式解析包内容 在外部解压后访问
func Load(pkg string, dst string) error {
	r, err := zip.OpenReader(pkg)
	if err != nil {
		return err
	}
	defer r.Close()
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}
	for _, f := range r.File {
		p := filepath.Join(dst, filepath.Clean(f.Name))
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(p, f.Mode()); err != nil {
				return err
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(p), f.Mode()); err != nil {
			return err
		}
		outFile, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		if _, err := io.Copy(outFile, rc); err != nil {
			return err
		}
		outFile.Close()
	}
	return nil
}
