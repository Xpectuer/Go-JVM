package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZipEntry is entry of zip
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

// readClass read ZIP file of absPath
func (ze *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(ze.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, ze, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

// string is a toString() method
func (ze *ZipEntry) String() string {
	return ze.absPath
}
