package store

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/fallback template

//go:generate gowrap gen -d . -i Store -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/fallback -o fallback.go

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// StoreWithFallback implements Store interface wrapped with Prometheus metrics
type StoreWithFallback struct {
	implementations []Store
	interval        time.Duration
}

// NewStoreWithFallback takes several implementations of the Store and returns an instance of Store
// which calls all implementations concurrently with given interval and returns first non-error response.
func NewStoreWithFallback(interval time.Duration, impls ...Store) StoreWithFallback {
	return StoreWithFallback{implementations: impls, interval: interval}
}

// Close implements Store
func (_d StoreWithFallback) Close() (err error) {
	type _resultStruct struct {
		err error
	}
	var _res _resultStruct
	var _ch = make(chan _resultStruct, 0)
	var _errorsList []string
	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	for _i := 0; _i < len(_d.implementations); _i++ {
		go func(_impl Store) {
			err := _impl.Close()
			if err != nil {
				err = fmt.Errorf("%T: %v", _impl, err)
			}

			_ch <- _resultStruct{err}
		}(_d.implementations[_i])
		select {
		case _res = <-_ch:
			if _res.err == nil {
				return _res.err
			}
			_errorsList = append(_errorsList, _res.err.Error())
		case <-_ticker.C:
			_errorsList = append(_errorsList, fmt.Sprintf("%T: timeout", _d.implementations[_i]))

		}
	}
	err = fmt.Errorf(strings.Join(_errorsList, ";"))
	return
}

// Delete implements Store
func (_d StoreWithFallback) Delete(key string, opts ...DeleteOption) (err error) {
	type _resultStruct struct {
		err error
	}
	var _res _resultStruct
	var _ch = make(chan _resultStruct, 0)
	var _errorsList []string
	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	for _i := 0; _i < len(_d.implementations); _i++ {
		go func(_impl Store) {
			err := _impl.Delete(key, opts...)
			if err != nil {
				err = fmt.Errorf("%T: %v", _impl, err)
			}

			_ch <- _resultStruct{err}
		}(_d.implementations[_i])
		select {
		case _res = <-_ch:
			if _res.err == nil {
				return _res.err
			}
			_errorsList = append(_errorsList, _res.err.Error())
		case <-_ticker.C:
			_errorsList = append(_errorsList, fmt.Sprintf("%T: timeout", _d.implementations[_i]))

		}
	}
	err = fmt.Errorf(strings.Join(_errorsList, ";"))
	return
}

// Download implements Store
func (_d StoreWithFallback) Download(target string, key string, opts ...DownloadOption) (err error) {
	type _resultStruct struct {
		err error
	}
	var _res _resultStruct
	var _ch = make(chan _resultStruct, 0)
	var _errorsList []string
	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	for _i := 0; _i < len(_d.implementations); _i++ {
		go func(_impl Store) {
			err := _impl.Download(target, key, opts...)
			if err != nil {
				err = fmt.Errorf("%T: %v", _impl, err)
			}

			_ch <- _resultStruct{err}
		}(_d.implementations[_i])
		select {
		case _res = <-_ch:
			if _res.err == nil {
				return _res.err
			}
			_errorsList = append(_errorsList, _res.err.Error())
		case <-_ticker.C:
			_errorsList = append(_errorsList, fmt.Sprintf("%T: timeout", _d.implementations[_i]))

		}
	}
	err = fmt.Errorf(strings.Join(_errorsList, ";"))
	return
}

// DownloadTo implements Store
func (_d StoreWithFallback) DownloadTo(writer io.WriterAt, key string, opts ...DownloadOption) (err error) {
	type _resultStruct struct {
		err error
	}
	var _res _resultStruct
	var _ch = make(chan _resultStruct, 0)
	var _errorsList []string
	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	for _i := 0; _i < len(_d.implementations); _i++ {
		go func(_impl Store) {
			err := _impl.DownloadTo(writer, key, opts...)
			if err != nil {
				err = fmt.Errorf("%T: %v", _impl, err)
			}

			_ch <- _resultStruct{err}
		}(_d.implementations[_i])
		select {
		case _res = <-_ch:
			if _res.err == nil {
				return _res.err
			}
			_errorsList = append(_errorsList, _res.err.Error())
		case <-_ticker.C:
			_errorsList = append(_errorsList, fmt.Sprintf("%T: timeout", _d.implementations[_i]))

		}
	}
	err = fmt.Errorf(strings.Join(_errorsList, ";"))
	return
}

// Get implements Store
func (_d StoreWithFallback) Get(key string, opts ...GetOption) (ba1 []byte, err error) {
	type _resultStruct struct {
		ba1 []byte
		err error
	}
	var _res _resultStruct
	var _ch = make(chan _resultStruct, 0)
	var _errorsList []string
	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	for _i := 0; _i < len(_d.implementations); _i++ {
		go func(_impl Store) {
			ba1, err := _impl.Get(key, opts...)
			if err != nil {
				err = fmt.Errorf("%T: %v", _impl, err)
			}

			_ch <- _resultStruct{ba1, err}
		}(_d.implementations[_i])
		select {
		case _res = <-_ch:
			if _res.err == nil {
				return _res.ba1, _res.err
			}
			_errorsList = append(_errorsList, _res.err.Error())
		case <-_ticker.C:
			_errorsList = append(_errorsList, fmt.Sprintf("%T: timeout", _d.implementations[_i]))

		}
	}
	err = fmt.Errorf(strings.Join(_errorsList, ";"))
	return
}

// GetReader implements Store
func (_d StoreWithFallback) GetReader(key0 string, opts ...GetOption) (r1 io.ReadCloser, err error) {
	type _resultStruct struct {
		r1  io.ReadCloser
		err error
	}
	var _res _resultStruct
	var _ch = make(chan _resultStruct, 0)
	var _errorsList []string
	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	for _i := 0; _i < len(_d.implementations); _i++ {
		go func(_impl Store) {
			r1, err := _impl.GetReader(key0, opts...)
			if err != nil {
				err = fmt.Errorf("%T: %v", _impl, err)
			}

			_ch <- _resultStruct{r1, err}
		}(_d.implementations[_i])
		select {
		case _res = <-_ch:
			if _res.err == nil {
				return _res.r1, _res.err
			}
			_errorsList = append(_errorsList, _res.err.Error())
		case <-_ticker.C:
			_errorsList = append(_errorsList, fmt.Sprintf("%T: timeout", _d.implementations[_i]))

		}
	}
	err = fmt.Errorf(strings.Join(_errorsList, ";"))
	return
}

// List implements Store
func (_d StoreWithFallback) List(opts ...ListOption) (sa1 []string, err error) {
	type _resultStruct struct {
		sa1 []string
		err error
	}
	var _res _resultStruct
	var _ch = make(chan _resultStruct, 0)
	var _errorsList []string
	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	for _i := 0; _i < len(_d.implementations); _i++ {
		go func(_impl Store) {
			sa1, err := _impl.List(opts...)
			if err != nil {
				err = fmt.Errorf("%T: %v", _impl, err)
			}

			_ch <- _resultStruct{sa1, err}
		}(_d.implementations[_i])
		select {
		case _res = <-_ch:
			if _res.err == nil {
				return _res.sa1, _res.err
			}
			_errorsList = append(_errorsList, _res.err.Error())
		case <-_ticker.C:
			_errorsList = append(_errorsList, fmt.Sprintf("%T: timeout", _d.implementations[_i]))

		}
	}
	err = fmt.Errorf(strings.Join(_errorsList, ";"))
	return
}

// Name implements Store
func (_d StoreWithFallback) Name() (s1 string) {
	type _resultStruct struct {
		s1 string
	}
	var _res _resultStruct
	var _ch = make(chan _resultStruct, 0)

	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	for _i := 0; _i < len(_d.implementations); _i++ {
		go func(_impl Store) {
			s1 := _impl.Name()
			_ch <- _resultStruct{s1}
		}(_d.implementations[_i])
		select {
		case _res = <-_ch:
			return _res.s1
		case <-_ticker.C:
		}
	}

	return
}

// Options implements Store
func (_d StoreWithFallback) Options() (o1 Options) {
	type _resultStruct struct {
		o1 Options
	}
	var _res _resultStruct
	var _ch = make(chan _resultStruct, 0)

	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	for _i := 0; _i < len(_d.implementations); _i++ {
		go func(_impl Store) {
			o1 := _impl.Options()
			_ch <- _resultStruct{o1}
		}(_d.implementations[_i])
		select {
		case _res = <-_ch:
			return _res.o1
		case <-_ticker.C:
		}
	}

	return
}

// Upload implements Store
func (_d StoreWithFallback) Upload(path string, key string, opts ...UploadOption) (s1 string, err error) {
	type _resultStruct struct {
		s1  string
		err error
	}
	var _res _resultStruct
	var _ch = make(chan _resultStruct, 0)
	var _errorsList []string
	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	for _i := 0; _i < len(_d.implementations); _i++ {
		go func(_impl Store) {
			s1, err := _impl.Upload(path, key, opts...)
			if err != nil {
				err = fmt.Errorf("%T: %v", _impl, err)
			}

			_ch <- _resultStruct{s1, err}
		}(_d.implementations[_i])
		select {
		case _res = <-_ch:
			if _res.err == nil {
				return _res.s1, _res.err
			}
			_errorsList = append(_errorsList, _res.err.Error())
		case <-_ticker.C:
			_errorsList = append(_errorsList, fmt.Sprintf("%T: timeout", _d.implementations[_i]))

		}
	}
	err = fmt.Errorf(strings.Join(_errorsList, ";"))
	return
}

// UploadFrom implements Store
func (_d StoreWithFallback) UploadFrom(reader io.Reader, key string, opts ...UploadOption) (s1 string, err error) {
	type _resultStruct struct {
		s1  string
		err error
	}
	var _res _resultStruct
	var _ch = make(chan _resultStruct, 0)
	var _errorsList []string
	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	for _i := 0; _i < len(_d.implementations); _i++ {
		go func(_impl Store) {
			s1, err := _impl.UploadFrom(reader, key, opts...)
			if err != nil {
				err = fmt.Errorf("%T: %v", _impl, err)
			}

			_ch <- _resultStruct{s1, err}
		}(_d.implementations[_i])
		select {
		case _res = <-_ch:
			if _res.err == nil {
				return _res.s1, _res.err
			}
			_errorsList = append(_errorsList, _res.err.Error())
		case <-_ticker.C:
			_errorsList = append(_errorsList, fmt.Sprintf("%T: timeout", _d.implementations[_i]))

		}
	}
	err = fmt.Errorf(strings.Join(_errorsList, ";"))
	return
}
