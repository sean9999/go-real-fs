package realfs

import (
	"io"
	"io/fs"
	"os"
)

// WritableFs is a writable filesystem
type WritableFs interface {
	RealFS
	WriteFile(string, []byte, fs.FileMode) error
	OpenFile(string, int, fs.FileMode) (WritableFile, error)
	Remove(string) error
}

type writableFs struct {
	realFS
}

var _ WritableFs = (*writableFs)(nil)

type WritableFile interface {
	fs.File
	io.Writer
}

type writableFile struct {
	*os.File
}

func (wfs writableFs) OpenFile(incorrectPath string, flag int, mode fs.FileMode) (WritableFile, error) {
	path, err := wfs.correctPath(incorrectPath)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(path, flag, mode)
	if err != nil {
		return nil, err
	}
	return writableFile{f}, nil
}

func (wfs writableFs) WriteFile(incorrectPath string, data []byte, perm fs.FileMode) error {
	path, err := wfs.correctPath(incorrectPath)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, perm)
}

func (wfs writableFs) Remove(incorrectPath string) error {
	path, err := wfs.correctPath(incorrectPath)
	if err != nil {
		return err
	}
	return os.Remove(path)
}

func NewWritable() writableFs {
	wfs := writableFs{
		realFS{
			dirFs: os.DirFS("/"),
		},
	}
	return wfs
}
