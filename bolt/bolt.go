package bolt

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Unknwon/com"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/rai-project/store"
)

type Handler struct {
	db     *bolt.DB
	opts   *store.Options
	isOpen bool
}

func New(iopts ...store.Option) (store.Store, error) {
	opts := NewOptions()

	for _, o := range iopts {
		o(opts)
	}

	basePath, ok := opts.Context.Value(basePathKey).(string)
	if !ok || basePath == "" {
		return nil, errors.New("base path was not set")
	}
	if com.IsDir(basePath) {
		basePath = filepath.Join(basePath, "bolt.db")
	}

	bucket := opts.Bucket
	if bucket == "" {
		return nil, errors.New("bolt bucket was not set")
	}

	db, err := bolt.Open(basePath, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &Handler{
		db:     db,
		opts:   opts,
		isOpen: true,
	}, nil
}

func (h *Handler) Name() string {
	return "Bolt"
}

func (h *Handler) Close() error {
	if !h.isOpen {
		return nil
	}
	defer func() {
		h.isOpen = false
	}()
	return h.db.Close()
}

func (h *Handler) Delete(key string, d ...store.DeleteOption) error {
	if !h.isOpen {
		return errors.New("bolt database is not open")
	}
	db := h.db
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := []byte(h.opts.Bucket)
		b := tx.Bucket(bucket)
		err := b.Delete([]byte(key))
		return err
	})
	return err
}

func (h *Handler) Download(s string, s1 string, d ...store.DownloadOption) error {
	if !h.isOpen {
		return errors.New("bolt database is not open")
	}
	db := h.db
	_ = db
	panic(errors.New("*Handler.Download not implemented"))
}

func (h *Handler) DownloadTo(w io.WriterAt, s string, d ...store.DownloadOption) error {
	if !h.isOpen {
		return errors.New("bolt database is not open")
	}
	db := h.db
	_ = db
	panic(errors.New("*Handler.DownloadTo not implemented"))
}

func (h *Handler) Get(key string, g ...store.GetOption) ([]byte, error) {
	if !h.isOpen {
		return nil, errors.New("bolt database is not open")
	}
	db := h.db
	var value []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := []byte(h.opts.Bucket)
		b := tx.Bucket(bucket)
		v := b.Get([]byte(key))
		if v != nil {
			value = append(value, v...)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (h *Handler) GetReader(key string, g ...store.GetOption) (io.ReadCloser, error) {
	if !h.isOpen {
		return nil, errors.New("bolt database is not open")
	}
	buf, err := h.Get(key, g...)
	if err != nil {
		return nil, err
	}
	return ioutil.NopCloser(bytes.NewBuffer(buf)), nil
}

func (h *Handler) List(l ...store.ListOption) ([]string, error) {
	if !h.isOpen {
		return nil, errors.New("bolt database is not open")
	}
	db := h.db
	var keys []string

	err := db.View(func(tx *bolt.Tx) error {
		bucket := []byte(h.opts.Bucket)
		b := tx.Bucket(bucket)
		b.ForEach(func(k, v []byte) error {
			// Due to Byte slices returned from Bolt are only valid during a transaction.
			// Once the transaction has been committed or rolled back then the memory
			// they point to can be reused by a new page or can be unmapped from
			// virtual memory and you'll see an unexpected fault address panic when accessing it.
			// We copy the slice to retain it
			dst := make([]byte, len(k))
			copy(dst, k)

			keys = append(keys, string(dst))
			return nil
		})
		return nil
	})
	if err != nil {
		return []string{}, err
	}

	return keys, nil
}

func (h *Handler) Options() store.Options {
	return *h.opts
}

func (h *Handler) Upload(path string, key string, u ...store.UploadOption) (string, error) {
	if !h.isOpen {
		return "", errors.New("bolt database is not open")
	}

	file, err := os.Open(path)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read file during bolt upload.")
	}
	defer file.Close()

	return h.UploadFrom(file, key, u...)
}

func (h *Handler) UploadFrom(r io.Reader, key string, u ...store.UploadOption) (string, error) {
	if !h.isOpen {
		return "", errors.New("bolt database is not open")
	}
	db := h.db
	err := db.Update(func(tx *bolt.Tx) error {
		var err error
		bucket := []byte(h.opts.Bucket)
		buf := make([]byte, 32*1024)
		b := tx.Bucket(bucket)
		for {
			nr, er := r.Read(buf)
			if nr > 0 {
				err = b.Put([]byte(key), buf[0:nr])
				if err != nil {
					break
				}
			}
			if er != nil {
				if er != io.EOF {
					err = er
				}
				break
			}
		}
		return err
	})

	return key, err
}
