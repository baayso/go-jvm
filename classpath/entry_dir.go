package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// 目录形式的类路径
type DirEntry struct {
	// 目录的绝对路径
	absDir string
}

func newDirEntry(path string) *DirEntry {
	// 把参数转换成绝对路径
	absDir, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return &DirEntry{absDir}
}

// 先把目录和class文件名拼成一个完整的路径，然后调用 ioutil包提供的 ReadFile()函数读取class文件内容
func (d *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(d.absDir, className)
	data, err := ioutil.ReadFile(fileName)

	return data, d, err
}

func (d *DirEntry) String() string {
	return d.absDir
}
