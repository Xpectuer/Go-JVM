package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry is the normal directory entry
type DirEntry struct {
	absDir string
}

// newDirEntry returns a new directory entry instance
// use Abs() to offer support for relative path
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

// readClass returns a file content
func (de *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(de.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, de, err
}

// String is a toString method
func (de *DirEntry) String() string {
	return de.absDir
}
