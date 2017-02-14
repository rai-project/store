package store

import "io"

type Store interface {
	Options() Options
	Upload(path string, key string, opts ...UploadOption) error
	UploadFrom(reader io.Reader, key string, opts ...UploadOption) error
	Download(target string, key string, opts ...DownloadOption) error
	DownloadTo(writer io.WriterAt, key string, opts ...DownloadOption) error
	List(opts ...ListOption) ([]string, error)
	Name() string
}
