package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	Stirng() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.Contains(path, "*") {
		return newWildCardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
}
