package s3

import (
	"context"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"github.com/rai-project/store"
)

func (s *s3Client) Download(target string, key string, opts ...store.DownloadOption) error {
	file, err := os.Create(target)
	if err != nil {
		return errors.Wrap(err, "failed to create target file for s3 download.")
	}
	defer file.Close()
	return s.DownloadTo(file, key, opts...)
}

func (s *s3Client) DownloadTo(writer io.WriterAt, key0 string, opts ...store.DownloadOption) error {

	options := store.DownloadOptions{
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	key0 = strings.TrimLeft(key0, s.opts.BaseURL)
	key, err := url.QueryUnescape(key0)
	if err != nil {
		log.WithField("key", key0).Error("Failed to unescape ", key0)
		key = key0
	}

	_, err = s.downloader.Download(
		writer,
		&s3.GetObjectInput{
			Bucket: aws.String(s.opts.Bucket),
			Key:    aws.String(key),
		},
	)
	if err != nil {
		return errors.Wrapf(err, "Failed to get data from %s/%s\n", s.opts.Bucket, key)
	}
	return nil

}
