package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Unknwon/com"
	"github.com/pkg/errors"
)

type LocalFileSystem struct {
}

func NewLocalFileSystem() *LocalFileSystem {
	return &LocalFileSystem{}
}

func (l *LocalFileSystem) IsFile(loc *FileLocation) bool {
	return com.IsFile(loc.Path)
}

func (l *LocalFileSystem) IsDir(loc *FileLocation) bool {
	return com.IsDir(loc.Path)
}

func (l *LocalFileSystem) List(loc *FileLocation) ([]*FileLocation, error) {
	if !com.IsDir(loc.Path) {
		return nil, errors.Errorf("the %v specified is not a directory to be listed", loc.Path)
	}

	files, err := ioutil.ReadDir(loc.Path)
	if err != nil {
		return nil, err
	}

	fileLocations := make([]*FileLocation, len(files))
	for ii, file := range files {
		fileLocations[ii] = NewFileLocation(filepath.Join(loc.Path, file.Name()))
	}
	return fileLocations, nil
}

func (l *LocalFileSystem) Open(loc *FileLocation) (File, error) {
	return os.Open(loc.Path)
}

func (l *LocalFileSystem) Accept(loc *FileLocation) bool {
	if strings.HasPrefix(loc.Path, "file://") {
		return true
	}
	if strings.HasPrefix(loc.Path, "mem://") ||
		strings.HasPrefix(loc.Path, "s3://") ||
		strings.HasPrefix(loc.Path, "minio://") {
		return false
	}
	return true
}
