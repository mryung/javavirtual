package classpath

import "os"
import "fmt"
import "path/filepath"

type ClassPath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *ClassPath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	fmt.Printf("jre Dir %s\n", jreDir)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	fmt.Printf("jreLibPath Dir %s\n", jreLibPath)
	self.bootClasspath = newWildCardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	fmt.Printf("jreExtPath Dir %s\n", jreExtPath)
	self.extClasspath = newWildCardEntry(jreExtPath)
}

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

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *ClassPath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	fmt.Printf(">>>>>>>>>>>>>>>>>\n")
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *ClassPath) String() string {
	return self.userClasspath.String()
}
