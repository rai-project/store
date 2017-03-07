package s3

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rai-project/store"
)

var (
	DefaultListMax = 100
)

func (s *s3Client) List(opts ...store.ListOption) ([]string, error) {
	options := store.ListOptions{
		Max:     int64(DefaultListMax),
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	objs, err := s.client.ListObjects(&s3.ListObjectsInput{
		Bucket:  aws.String(s.opts.Bucket),
		MaxKeys: aws.Int64(options.Max),
	})
	if err != nil {
		return nil, err
	}
	keys := make([]string, len(objs.Contents))
	for ii, c := range objs.Contents {
		keys[ii] = aws.StringValue(c.Key)
	}
	return keys, nil
}
