package fs

import (
	"io"

	"github.com/pkg/errors"
	"github.com/rai-project/aws"
	"github.com/rai-project/config"
	"github.com/rai-project/store/s3"
)

type File interface {
	io.ReadCloser
}

type FileSystem interface {
	Accept(*FileLocation) bool
	Open(*FileLocation) (File, error)
	List(*FileLocation) ([]*FileLocation, error)
	IsFile(*FileLocation) bool
	IsDir(*FileLocation) bool
}

type FileLocation struct {
	Path string
}

var (
	fileSystems []FileSystem
)

func init() {
	config.AfterInit(func() {
		aws.Config.Wait()
		s3.Config.Wait()
		fileSystems = []FileSystem{
			NewLocalFileSystem(),
			NewMemoryFileSystem(),
			NewS3FileSystem(),
		}
	})
}

func NewFileLocation(filePath string) *FileLocation {
	return &FileLocation{
		Path: filePath,
	}
}

func findFileSystem(loc *FileLocation) (FileSystem, error) {
	for _, vfs := range fileSystems {
		if vfs.Accept(loc) {
			return vfs, nil
		}
	}
	return nil, errors.Errorf("unable to handle %v file path", loc.Path)
}

func Open(filePath string) (File, error) {
	loc := NewFileLocation(filePath)
	vfs, err := findFileSystem(loc)
	if err != nil {
		return nil, err
	}
	return vfs.Open(loc)
}

func List(filePath string) ([]*FileLocation, error) {
	loc := NewFileLocation(filePath)
	vfs, err := findFileSystem(loc)
	if err != nil {
		return nil, err
	}
	return vfs.List(loc)
}

func IsFile(filePath string) bool {
	loc := NewFileLocation(filePath)
	vfs, err := findFileSystem(loc)
	if err != nil {
		return false
	}
	return vfs.IsFile(loc)
}

func IsDir(filePath string) bool {
	loc := NewFileLocation(filePath)
	vfs, err := findFileSystem(loc)
	if err != nil {
		return false
	}
	return vfs.IsDir(loc)
}
