package s3

import (
	"context"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"github.com/rai-project/store"
)

func (s *s3Client) Get(key0 string, opts ...store.GetOption) ([]byte, error) {

	options := store.GetOptions{
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

	obj, err := s.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.opts.Bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return []byte{}, errors.Wrapf(err, "Failed to get data from %s/%s\n", s.opts.Bucket, key)
	}
	if obj != nil {
		defer obj.Body.Close()
	}
	ret, err := ioutil.ReadAll(obj.Body)
	//  meta := obj.Metadata
	return ret, nil
}
