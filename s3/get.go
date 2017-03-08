package s3

import (
	"context"
	"io"
	"io/ioutil"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"github.com/rai-project/store"
)

func (s *s3Client) Get(key0 string, opts ...store.GetOption) ([]byte, error) {
	body, err := s.GetReader(key0, opts...)
	if err != nil {
		return nil, err
	}
	if body != nil {
		defer body.Close()
	}
	ret, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *s3Client) GetReader(key0 string, opts ...store.GetOption) (io.ReadCloser, error) {

	options := store.GetOptions{
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	prefix := s.opts.BaseURL + "/" + s.opts.Bucket + "/"
	key0 = cleanupKey(key0, prefix)

	key, err := url.QueryUnescape(key0)
	if err != nil {
		log.WithField("key", key0).Error("Failed to unescape ", key0)
		key = key0
	}

	obj, err := s.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.opts.Bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get data from %s/%s\n", s.opts.Bucket, key)
	}

	return obj.Body, nil
}
