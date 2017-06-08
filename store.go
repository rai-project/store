package store

import (
	"io"
)

type Store interface {
	Options() Options
	Upload(path string, key string, opts ...UploadOption) (string, error)
	UploadFrom(reader io.Reader, key string, opts ...UploadOption) (string, error)
	Download(target string, key string, opts ...DownloadOption) error
	DownloadTo(writer io.WriterAt, key string, opts ...DownloadOption) error
	Get(key string, opts ...GetOption) ([]byte, error)
	GetReader(key0 string, opts ...GetOption) (io.ReadCloser, error)
	List(opts ...ListOption) ([]string, error)
	Delete(key string, opts ...DeleteOption) error
	Name() string
	Close() error
}
