package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

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

func (self *ZipEntry) String() string {
	return self.absPath
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)

	if err != nil {
		return nil, nil, err
	}

	defer r.Close()
	for _, f := rang r.File {
		if f.Name == className {
			rc,err := f.Open()
			if err != nil {
				return nil,nil,err
			}
			defer rc.Close()
			data,err := ioutil.ReadFile(rc)
			if err != nil{
				return nil,nil,err
			}

			return data,self,nil
		}
	}
}