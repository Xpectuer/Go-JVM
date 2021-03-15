package classpath

import (
	"os"
	"strings"
)

// 实现类路径implements classpath
/**
class path contains:
1. bootstrap classpath
2. extension classpath
3. user classpath
The bootstrap classpath can be mapped to "jre/lib" directory by the default,
which contains the most of the Java Standard Library.

The extention classpath can be mapped to "jre/lib/ext" directory,
which contains the classes that employ Java Extension Manchanism.

The user classpath by the default use the current directory which is "."
The CLASSPATH Environment variable can be set as the user classpath,
but this is not recommended for its inflexibility.
A better way to pass the classpath is to set parameter using -classpath (abbreviate for -cp) option.
(The -classpath/-cp option has a higher priority than CLASSPATH environment variable)

example:
java -cp path/to/classes;lib/a.jar;lib/b.jar;lib/c.zip ...
since java 6
java -cp classes;lib/* ...
*/

// pathListSeperator存放类路径分隔符
const pathListSeperator = string(os.PathListSeparator)

// Entry is the classpath entry
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

// newEntry initialized an Entry
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeperator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)

}
