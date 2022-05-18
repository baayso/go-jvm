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

// 首先打开ZIP文件，如果这一步出错的话，直接返回。
// 然后遍历ZIP压缩包里的文件，看能否找到class文件。
// 如果能找到，则打开class文件，把内容读取出来，并返回。
// 如果找不到，或者出现其他错误，则返回错误信息。
// 有两处使用了defer语句来确保打开的文件得以关闭。
// readClass()方法每次都要打开和关闭ZIP文件，因此效率不是很高。
func (z *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(z.absPath)

	if err != nil {
		return nil, nil, err
	}

	defer r.Close()

	for _, f := range r.File {
		if f.Name == className {
			return func() ([]byte, Entry, error) {

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
			}()
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (z *ZipEntry) String() string {
	return z.absPath
}
