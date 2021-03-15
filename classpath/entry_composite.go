package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry is a entry that contains multiple entry
// to handle the situation like
// java -cp path/to/classes;lib/a.jar;lib/b.jar;lib/c.zip ...
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeperator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry

}

// readClass is  a implementation of readClass() method in the interface Entry
func (ce CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range ce {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (ce CompositeEntry) String() string {
	strs := make([]string, len(ce))
	for i, entry := range ce {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeperator)
}
