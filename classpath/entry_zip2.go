package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry2 struct {
	absPath string
	zipRC   *zip.ReadCloser
}

func newZipEntry2(path string) *ZipEntry2 {
	absPath, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return &ZipEntry2{absPath, nil}
}

func (z *ZipEntry2) readClass(className string) ([]byte, Entry, error) {
	if z.zipRC == nil {
		err := z.openJar()

		if err != nil {
			return nil, nil, err
		}
	}

	classFile := z.findClass(className)

	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)

	return data, z, err
}

// todo: close zip
func (z *ZipEntry2) openJar() error {
	r, err := zip.OpenReader(z.absPath)

	if err == nil {
		z.zipRC = r
	}

	return err
}

func (z *ZipEntry2) findClass(className string) *zip.File {
	for _, f := range z.zipRC.File {

		if f.Name == className {
			return f
		}
	}

	return nil
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()

	if err != nil {
		return nil, err
	}

	defer rc.Close()

	// read class data
	data, err := ioutil.ReadAll(rc)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (z *ZipEntry2) String() string {
	return z.absPath
}
