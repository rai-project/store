package fs

import (
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

type MemoryFileSystem struct {
	afero.Fs
}

func NewMemoryFileSystem() *MemoryFileSystem {
	return &MemoryFileSystem{Fs: afero.NewMemMapFs()}
}

func (l *MemoryFileSystem) IsFile(loc *FileLocation) bool {
	filePath := loc.Path
	f, e := l.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

func (l *MemoryFileSystem) IsDir(loc *FileLocation) bool {
	f, e := l.Stat(loc.Path)
	if e != nil {
		return false
	}
	return f.IsDir()
}

func (l *MemoryFileSystem) Open(loc *FileLocation) (File, error) {
	return l.Fs.Open(loc.Path)
}

func (l *MemoryFileSystem) List(loc *FileLocation) ([]*FileLocation, error) {
	if !l.IsDir(loc) {
		return nil, errors.Errorf("the %v specified is not a directory to be listed", loc.Path)
	}

	f, err := l.Fs.Open(loc.Path)
	if err != nil {
		return nil, err
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}

	fileLocations := make([]*FileLocation, len(files))
	for ii, file := range files {
		fileLocations[ii] = NewFileLocation(filepath.Join(loc.Path, file.Name()))
	}
	return fileLocations, nil
}

func (l *MemoryFileSystem) Accept(loc *FileLocation) bool {
	return strings.HasPrefix(loc.Path, "mem://")
}
