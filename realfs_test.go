package realfs

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRealFs struct {
	mock.Mock
}

func (mrfs *mockRealFs) ReadDir(path string) ([]fs.DirEntry, error) {
	args := mrfs.Called(path)
	return args.Get(0).([]fs.DirEntry), args.Error(1)
}

var myTestFs fstest.MapFS = fstest.MapFS{
	"Documents": &fstest.MapFile{
		Mode: fs.ModeDir,
	},
	"Downloads": &fstest.MapFile{
		Mode: fs.ModeDir,
	},
	".bashrc": &fstest.MapFile{
		Data: []byte("#!/bin/bash\n\nalias hello='echo hello'\n"),
	},
	"Downloads/meldmerge.dmg":     &fstest.MapFile{},
	"Downloads/JetBrainsMono.zip": &fstest.MapFile{},
}

func TestNewMockRealFs(t *testing.T) {

	m := new(mockRealFs)
	m.On("ReadDir", ".").Return(myTestFs.ReadDir("."))

	expected, err := myTestFs.ReadDir(".")
	if err != nil {
		t.Error(err)
	}
	actual, err := m.ReadDir(".")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, actual, "ah crap!")

}
