package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // remove *
	compositeEntry := []Entry{}
	// All errors that arise visiting files
	// and directories are filtered by walkFn.
	walkFn := func(path string, info os.FileInfo, err error) error {
		// return outer error
		if err != nil {
			return err
		}
		// skip if it is the Sub Directory
		// In linux file trees the walk function will see like this
		// . .. file1 dir1 file2 a b c
		// . means current directory
		// .. means parent directory
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
