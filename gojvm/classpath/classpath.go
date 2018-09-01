package classpath

import (
	"os"
	"path/filepath"
)

// 启动类路径默认对应jre\lib目录，Java标准库（大部分在rt.jar里）位于该路径。
// 扩展类路径默认对应jre\lib\ext目录，使用Java扩展机制的类位于这个路径。
// 我们自己实现的类，以及第三方类库则位于用户类路径。可以通过-Xbootclasspath选项修改启动类路径，不过通常并不需要这样做。
//
// 用户类路径的默认值是当前目录，也就是“.”。可以设置CLASSPATH环境变量来修改用户类路径，
// 但是这样做不够灵活，所以不推荐使用。更好的办法是给java命令传递-classpath（或简写为-cp）选项。
// -classpath/-cp选项的优先级更高，可以覆盖CLASSPATH环境变量设置。
// -classpath/-cp选项既可以指定目录，也可以指定JAR文件或者ZIP文件，还可以同时指定多个目录或文件，用分隔符分开即可。
//
// 分隔符因操作系统而异。在Windows系统下是分号，在类UNIX（包括Linux、Mac OS X等）系统下是冒号。例如在Windows下：
// java -cp path\to\classes;lib\a.jar;lib\b.jar;lib\c.zip ...
type Classpath struct {
	bootClasspath Entry // 启动类路径（bootstrap classpath）
	extClasspath  Entry // 扩展类路径（extension classpath）
	userClasspath Entry // 用户类路径（user classpath）
}

// 如果用户没有提供 -classpath/-cp 选项，则使用当前目录作为用户类路径。
// ReadClass()方法依次从启动类路径、扩展类路径和用户类路径中搜索class文件。
// 注意：传递给 ReadClass()方法的类名不包含“.class”后缀。
func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	if data, entry, err := c.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := c.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	return c.userClasspath.readClass(className)
}

// 返回用户类路径的字符串表示
func (c *Classpath) String() string {
	return c.userClasspath.String()
}

func (c *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClasspath = newWildcardEntry(jreExtPath)
}

func (c *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}

	c.userClasspath = newEntry(cpOption)
}

// 优先使用用户输入的 -Xjre选项作为jre目录。如果没有输入该选项，则在当前目录下寻找jre目录。
// 如果找不到，尝试使用 JAVA_HOME环境变量。
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre folder!")
}

// 用于判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// 使用 -Xjre 选项解析启动类路径和扩展类路径，使用 -classpath/-cp 选项解析用户类路径
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}

	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)

	return cp
}
