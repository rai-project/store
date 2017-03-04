package s3

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rai-project/store"
)

func (s *s3Client) Delete(key string, opts ...store.DeleteOption) error {
	options := store.ListOptions{
		Context: context.Background(),
	}
	for _, o := range opts {
		o(&options)
	}

	_, err := s.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.opts.Bucket),
		Key:    aws.String(key),
	})
	return err
}
