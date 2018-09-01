package classpath

import (
	"os"
	"strings"
)

// 路径分隔符
// :(linux/unix) or ;(windows)
const pathListSeparator = string(os.PathListSeparator)

// 定义一个接口来表示类路径项
type Entry interface {
	// 负责寻找和加载class文件
	// 入参(className)格式: java/lang/Object.class
	// 返回值：读取到的字节数据、最终定位到class文件的Entry，以及错误信息
	readClass(className string) ([]byte, Entry, error)

	// 返回字符串表示
	String() string
}

// 根据参数创建不同类型的 Entry实例
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
