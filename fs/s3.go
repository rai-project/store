package fs

import (
	"errors"
	"strings"

	"net/url"

	"path/filepath"

	"math"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/rai-project/aws"
	"github.com/rai-project/store"
	"github.com/rai-project/store/s3"
)

type S3FileSystem struct {
	session *session.Session
}

func newS3Client(sess *session.Session, iopts ...store.Option) (store.Store, error) {
	opts := append([]store.Option{s3.Session(sess)}, iopts...)
	client, err := s3.New(opts...)
	if err != nil {
		log.WithError(err).Fatal("unable to create an s3 client")
		return nil, errors.New("unable to create an s3 client")
	}
	return client, nil
}

func NewS3FileSystem() *S3FileSystem {
	sess, err := aws.NewSession()
	if err != nil {
		log.WithError(err).Fatal("unable to create an aws session")
		return nil
	}
	return &S3FileSystem{
		session: sess,
	}
}

func (l *S3FileSystem) IsFile(loc *FileLocation) bool {
	return false
}

func (l *S3FileSystem) IsDir(loc *FileLocation) bool {
	return false
}

func trimS3Prefix(s string) string {
	return strings.TrimLeft(
		strings.TrimLeft(s, "minio://"),
		"s3://",
	)
}

func (l *S3FileSystem) newS3ClientForLocation(loc *FileLocation) (string, store.Store, error) {
	// TODO ::: Double check this
	u, err := url.Parse(loc.Path)
	if err != nil {
		return "", nil, err
	}
	baseURL := u.Hostname()
	bucket := filepath.Dir(u.Path)
	file := filepath.Base(u.Path)
	s, err := newS3Client(
		l.session,
		store.BaseURL(baseURL),
		store.Bucket(bucket),
	)
	return file, s, err
}

func (l *S3FileSystem) Open(loc *FileLocation) (File, error) {
	key, s, err := l.newS3ClientForLocation(loc)
	if err != nil {
		return nil, err
	}
	return s.GetReader(key)
}

func (l *S3FileSystem) List(loc *FileLocation) ([]*FileLocation, error) {
	_, s, err := l.newS3ClientForLocation(loc)
	if err != nil {
		return nil, err
	}
	ls, err := s.List(store.Max(math.MaxInt64))
	if err != nil {
		return nil, err
	}
	res := make([]*FileLocation, len(ls))
	for ii, f := range ls {
		res[ii] = NewFileLocation(f)
	}
	return res, nil
}

func (l *S3FileSystem) Accept(loc *FileLocation) bool {
	return strings.HasPrefix(loc.Path, "s3://") ||
		strings.HasPrefix(loc.Path, "minio://")
}
