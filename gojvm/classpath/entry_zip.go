package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZIP或JAR文件形式的类路径
type ZipEntry struct {
	// 存放ZIP或JAR文件的绝对路径
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	// 把参数转换成绝对路径
	absPath, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath}
}

func (z *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(z.absPath)

	if err != nil {
		return nil, nil, err
	}

	defer r.Close()

	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()

			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()

			data, err := ioutil.ReadAll(rc)

			if err != nil {
				return nil, nil, err
			}

			return data, z, err
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (z *ZipEntry) String() string {
	return z.absPath
}
