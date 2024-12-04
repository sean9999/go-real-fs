package realfs

import (
	"io/fs"
	"testing/fstest"
)

type TestFS struct {
	mfs fstest.MapFS
}

func NewTestFs() TestFS {
	mfs := fstest.MapFS{}
	tfs := TestFS{mfs}
	return tfs
}

var _ WritableFs = (*TestFS)(nil)

func (tw TestFS) Open(name string) (fs.File, error) {
	k, exists := tw.mfs[name]
	if !exists {
		return nil, fs.ErrNotExist
	}

	fyle := TestFSFile{k, name, nil, 0}

	fi := &finfo{
		name, false, &fyle,
	}

	fyle.Info = fi

	return &fyle, nil
}

func (tw TestFS) WriteFile(name string, data []byte, mode fs.FileMode) error {

	tw.mfs[name] = &fstest.MapFile{
		Data: data,
		Mode: mode,
	}

	return nil
}

func (tw TestFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return nil, nil
}

func (tw TestFS) ReadFile(name string) ([]byte, error) {
	f, exists := tw.mfs[name]
	if !exists {
		return nil, fs.ErrNotExist
	}
	return f.Data, nil
}

func (tw TestFS) Remove(name string) error {
	delete(tw.mfs, name)
	return nil
}

func (tw TestFS) Stat(name string) (fs.FileInfo, error) {
	// thing, exists := tw.mfs[name]
	// if !exists {
	// 	return nil, fs.ErrNotExist
	// }
	thing, err := tw.Open(name)
	if err != nil {
		return nil, err
	}
	return thing.Stat()
}

func (tw TestFS) OpenFile(name string, _ int, _ fs.FileMode) (WritableFile, error) {

	f, err := tw.Open(name)
	if err != nil {
		return nil, err
	}

	return f.(WritableFile), nil
}
