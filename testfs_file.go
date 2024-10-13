package realfs

import (
	"io"
	"io/fs"
	"regexp"
	"testing/fstest"
	"time"
)

type finfo struct {
	name  string
	isDir bool
	f     *TestFSFile
}

var _ fs.File = (*TestFSFile)(nil)

type TestFSFile struct {
	*fstest.MapFile
	name   string
	Info   *finfo
	cursor int64
}

func (f *TestFSFile) IsDir() bool {
	//	yes if the last char is "/"
	reg := regexp.MustCompile("\\/$")
	return reg.MatchString(f.name)
}

func (f *finfo) IsDir() bool {
	return f.f.IsDir()
}

func (f *finfo) ModTime() time.Time {
	return f.f.ModTime
}

func (_ *TestFSFile) Close() error {
	return nil
}

func (f *TestFSFile) Name() string {
	return f.name
}

func (f *TestFSFile) Seek(offest int64, whence int) (int64, error) {

	var startingPoint, newCursor int64

	switch whence {
	case io.SeekStart:
		startingPoint = 0
	case io.SeekCurrent:
		startingPoint = f.cursor
	case io.SeekEnd:
		startingPoint = int64(len(f.Data)) - 1
	}

	newCursor = startingPoint + offest

	if newCursor < 0 || newCursor >= int64(len(f.Data)) {
		return 0, io.ErrNoProgress
	}

	f.cursor = newCursor
	return newCursor, nil
}

func (f *TestFSFile) Read(p []byte) (int, error) {
	unreadData := f.Data[f.cursor:]
	if len(unreadData) == 0 {
		return 0, io.EOF
	}
	bytesWritten := copy(p, unreadData)
	f.cursor += int64(bytesWritten)
	return bytesWritten, nil
}

func (f *TestFSFile) Stat() (fs.FileInfo, error) {
	return f.Info, nil
}

func (f *finfo) Name() string {
	return f.f.name
}

func (f *finfo) Mode() fs.FileMode {
	return f.f.Mode
}

func (f *finfo) Size() int64 {
	return int64(len(f.f.Data))
}

func (f *finfo) Sys() any {
	return map[string]string{
		"type": "testfs",
	}
}
