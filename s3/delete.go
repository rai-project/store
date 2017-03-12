package s3

import (
	"context"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rai-project/store"
)

func (s *s3Client) Delete(key0 string, opts ...store.DeleteOption) error {
	options := store.ListOptions{
		Context: context.Background(),
	}
	for _, o := range opts {
		o(&options)
	}

	prefix := s.opts.BaseURL + "/" + s.opts.Bucket + "/"
	key0 = strings.TrimPrefix(
		strings.TrimPrefix(
			strings.TrimPrefix(
				key0,
				"http://",
			),
			"https://",
		),
		prefix,
	)

	key, err := url.QueryUnescape(key0)
	if err != nil {
		key = key0
	}

	_, err = s.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.opts.Bucket),
		Key:    aws.String(key),
	})

	return err
}
