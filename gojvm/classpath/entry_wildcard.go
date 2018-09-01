package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// 首先把路径末尾的星号去掉，得到 baseDir，然后调用 filepath包的 Walk()函数遍历 baseDir创建 ZipEntry
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // 把路径末尾的星号去掉

	compositeEntry := CompositeEntry{}

	// 在 walkFn中，根据后缀名选出 jar文件，并且返回 SkipDir跳过子目录（通配符类路径不能递归匹配子目录下的jar文件）。
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}

		return nil
	}

	filepath.Walk(baseDir, walkFn)

	return compositeEntry
}
