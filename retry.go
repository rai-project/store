package store

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/retry template

//go:generate gowrap gen -d . -i Store -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/retry -o retry.go

import (
	"io"
	"time"
)

// StoreWithRetry implements Store interface instrumented with retries
type StoreWithRetry struct {
	Store
	_retryCount    int
	_retryInterval time.Duration
}

// NewStoreWithRetry returns StoreWithRetry
func NewStoreWithRetry(base Store, retryCount int, retryInterval time.Duration) StoreWithRetry {
	return StoreWithRetry{
		Store:          base,
		_retryCount:    retryCount + 1,
		_retryInterval: retryInterval,
	}
}

// Close implements Store
func (_d StoreWithRetry) Close() (err error) {
	for _i := 0; _i < _d._retryCount; _i++ {
		err = _d.Store.Close()
		if err == nil {
			break
		}
		if _d._retryCount > 1 {
			time.Sleep(_d._retryInterval)
		}
	}
	return err
}

// Delete implements Store
func (_d StoreWithRetry) Delete(key string, opts ...DeleteOption) (err error) {
	for _i := 0; _i < _d._retryCount; _i++ {
		err = _d.Store.Delete(key, opts...)
		if err == nil {
			break
		}
		if _d._retryCount > 1 {
			time.Sleep(_d._retryInterval)
		}
	}
	return err
}

// Download implements Store
func (_d StoreWithRetry) Download(target string, key string, opts ...DownloadOption) (err error) {
	for _i := 0; _i < _d._retryCount; _i++ {
		err = _d.Store.Download(target, key, opts...)
		if err == nil {
			break
		}
		if _d._retryCount > 1 {
			time.Sleep(_d._retryInterval)
		}
	}
	return err
}

// DownloadTo implements Store
func (_d StoreWithRetry) DownloadTo(writer io.WriterAt, key string, opts ...DownloadOption) (err error) {
	for _i := 0; _i < _d._retryCount; _i++ {
		err = _d.Store.DownloadTo(writer, key, opts...)
		if err == nil {
			break
		}
		if _d._retryCount > 1 {
			time.Sleep(_d._retryInterval)
		}
	}
	return err
}

// Get implements Store
func (_d StoreWithRetry) Get(key string, opts ...GetOption) (ba1 []byte, err error) {
	for _i := 0; _i < _d._retryCount; _i++ {
		ba1, err = _d.Store.Get(key, opts...)
		if err == nil {
			break
		}
		if _d._retryCount > 1 {
			time.Sleep(_d._retryInterval)
		}
	}
	return ba1, err
}

// GetReader implements Store
func (_d StoreWithRetry) GetReader(key0 string, opts ...GetOption) (r1 io.ReadCloser, err error) {
	for _i := 0; _i < _d._retryCount; _i++ {
		r1, err = _d.Store.GetReader(key0, opts...)
		if err == nil {
			break
		}
		if _d._retryCount > 1 {
			time.Sleep(_d._retryInterval)
		}
	}
	return r1, err
}

// List implements Store
func (_d StoreWithRetry) List(opts ...ListOption) (sa1 []string, err error) {
	for _i := 0; _i < _d._retryCount; _i++ {
		sa1, err = _d.Store.List(opts...)
		if err == nil {
			break
		}
		if _d._retryCount > 1 {
			time.Sleep(_d._retryInterval)
		}
	}
	return sa1, err
}

// Upload implements Store
func (_d StoreWithRetry) Upload(path string, key string, opts ...UploadOption) (s1 string, err error) {
	for _i := 0; _i < _d._retryCount; _i++ {
		s1, err = _d.Store.Upload(path, key, opts...)
		if err == nil {
			break
		}
		if _d._retryCount > 1 {
			time.Sleep(_d._retryInterval)
		}
	}
	return s1, err
}

// UploadFrom implements Store
func (_d StoreWithRetry) UploadFrom(reader io.Reader, key string, opts ...UploadOption) (s1 string, err error) {
	for _i := 0; _i < _d._retryCount; _i++ {
		s1, err = _d.Store.UploadFrom(reader, key, opts...)
		if err == nil {
			break
		}
		if _d._retryCount > 1 {
			time.Sleep(_d._retryInterval)
		}
	}
	return s1, err
}
