package classpath

import (
	"path/filepath"
	"os"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string)  *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string){
	jreDir := getGreDir(jreOption)
	jreLibPath := filepath.Join(jreOption, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func getGreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if js := os.Getenv("JAVA_HOME"); js != "" {
		return filepath.Join(js, "jre")
	}

	panic("Can noe find jre folder")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}


func (self *Classpath) parseUserClasspath(cpOption string){
	if cpOption == "" {
		cpOption = "."
	}

	self.userClasspath = newEntry(cpOption)
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	return self.userClasspath.readClass(className)

}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}