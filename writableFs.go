package realfs

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
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
	io.Seeker
	Name() string
}

type writableFile struct {
	*os.File
}

func (wfs writableFs) CreateFile(name string, data []byte) error {
	return errors.New("not implemented")
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
	// path, err := wfs.correctPath(incorrectPath)
	// if err != nil {
	// 	return err
	// }

	fullPath, err := filepath.Abs(incorrectPath)
	if err != nil {
		return err
	}

	return os.WriteFile(fullPath, data, perm)
}

func (wfs writableFs) Remove(incorrectPath string) error {
	// path, err := wfs.correctPath(incorrectPath)
	// if err != nil {
	// 	return err
	// }
	return os.Remove(incorrectPath)
}

func NewWritable() writableFs {
	wfs := writableFs{
		realFS{
			dirFs: os.DirFS("/"),
		},
	}
	return wfs
}
