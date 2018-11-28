package classpath

import (
	"path/filepath"
	"io/ioutil"
)


// 2.3.2
type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)

	if err !=nil {
		panic(err)
	}

	return  &DirEntry{absDir}
}

func (self *DirEntry) readsClass(className string) ([]byte, *DirEntry, error) {
	fileName := filepath.Join(self.absDir, className)

	data, err := ioutil.ReadFile(fileName)

	return  data, self, err
}

func (self *DirEntry) String() string {
	return  self.absDir
}





