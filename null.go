package realfs

import (
	"io"
	"io/fs"
)

type NullDevice struct {
	io.Writer
}

func (b NullDevice) Read(_ []byte) (int, error) {
	return 0, nil
}
func (b NullDevice) Open(_ string) (fs.File, error) {
	return nil, nil
}

func (b NullDevice) ReadDir(_ string) ([]fs.DirEntry, error) {
	return nil, nil
}

func (b NullDevice) ReadFile(_ string) ([]byte, error) {
	return nil, nil
}

func (b NullDevice) Stat(_ string) (fs.FileInfo, error) {
	return nil, nil
}

func (b NullDevice) OpenFile(name string, flag int, perm fs.FileMode) (WritableFile, error) {
	return nil, nil
}

func (b NullDevice) Remove(_ string) error {
	return nil
}

func (b NullDevice) WriteFile(_ string, _ []byte, _ fs.FileMode) error {
	return nil
}
