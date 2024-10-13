package realfs

import (
	"io/fs"
	"os"
	"path/filepath"
)

// RealFS is real filesystem that wraps [os.DirFS] and implements [fs.FS]
type RealFS interface {
	fs.StatFS
	fs.ReadFileFS
	fs.ReadDirFS
}

type realFS struct {
	dirFs fs.FS
}

// realFS implements RealFS
var _ RealFS = (*realFS)(nil)

// since an fs.FS expects file paths that don't work in the real world
// we transform to root-relative
// ex:	~/.bashrc	becomes	home/henry/.bashrc
// ex:	/etc/passwd	becomes	etc/passwd
func (rfs realFS) correctPath(relativePath string) (string, error) {
	fullPath, err := filepath.Abs(relativePath)
	// if err != nil {
	// 	return fullPath, err
	// }
	// return filepath.Rel("/", fullPath)
	return fullPath, err
}

func (rfs realFS) Open(name string) (fs.File, error) {
	newName, err := rfs.correctPath(name)
	if err != nil {
		return nil, err
	}
	return rfs.dirFs.Open(newName)
}

func (rfs realFS) Stat(name string) (fs.FileInfo, error) {
	newName, err := rfs.correctPath(name)
	if err != nil {
		return nil, err
	}
	return rfs.dirFs.(fs.StatFS).Stat(newName)
}

func (rfs realFS) ReadDir(name string) ([]fs.DirEntry, error) {
	newName, err := rfs.correctPath(name)
	if err != nil {
		return nil, err
	}
	return rfs.dirFs.(fs.ReadDirFS).ReadDir(newName)
}

func (rfs realFS) ReadFile(name string) ([]byte, error) {
	newName, err := rfs.correctPath(name)
	if err != nil {
		return nil, err
	}
	return rfs.dirFs.(fs.ReadFileFS).ReadFile(newName)
}

func New() RealFS {
	return realFS{
		dirFs: os.DirFS("/"),
	}
}
