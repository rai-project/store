package s3

import (
	"context"
	"io"
	"mime"
	"os"
	"time"

	"path/filepath"

	"github.com/Unknwon/com"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	humanize "github.com/dustin/go-humanize"
	"github.com/pkg/errors"
	"github.com/rai-project/store"
	"github.com/rai-project/uuid"
)

func (s *s3Client) createBucket(bucket string) error {
	buckets, err := s.client.ListBuckets(nil)
	if err != nil {
		log.WithError(err).Error("Cannot list buckets")
		return errors.Wrap(err, "cannot list buckets")
	}

	// try to find existing bucket
	for _, b := range buckets.Buckets {
		if aws.StringValue(b.Name) == bucket {
			return nil
		}
	}

	cparams := &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	}

	// Create a new bucket using the CreateBucket call.
	_, err = s.client.CreateBucket(cparams)
	if err != nil {
		// Message from an error.
		log.WithError(err).Error("Cannot create bucket " + bucket)
		return errors.Wrap(err, "cannot create bucket")
	}
	return nil
}

func (s *s3Client) Upload(path string, key string, opts ...store.UploadOption) (string, error) {

	if !com.IsFile(path) {
		return "", errors.Errorf("The file %v was not found", path)
	}

	ext := filepath.Ext(path)
	mimetype := mime.TypeByExtension(ext)

	options := store.UploadOptions{
		Context: context.WithValue(
			context.Background(),
			mimetypeKey,
			mimetype,
		),
	}

	for _, o := range opts {
		o(&options)
	}

	if fileSizeLimit, ok := options.Context.Value(fileSizeLimitKey).(int64); ok {
		fileSize, err := com.FileSize(path)
		if err != nil {
			return "", errors.Wrapf(err, "Cannot get file size for %s.", path)
		}
		if fileSize > fileSizeLimit {
			return "", errors.Errorf(
				"File size of %v exceeded limit of %v",
				humanize.Bytes(uint64(fileSize)),
				humanize.Bytes(uint64(fileSizeLimit)),
			)
		}
	}

	if key == "" {
		key = uuid.New(path)
	}

	file, err := os.Open(path)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read file during s3 upload.")
	}
	defer file.Close()

	return s.UploadFrom(file, key, opts...)
}

func (s *s3Client) UploadFrom(reader io.Reader, key string, opts ...store.UploadOption) (string, error) {

	options := store.UploadOptions{
		Context: context.WithValue(
			context.Background(),
			aclKey,
			Config.ACL,
		),
	}

	for _, o := range opts {
		o(&options)
	}

	if err := s.createBucket(s.opts.Bucket); err != nil {
		return "", err
	}

	if key == "" {
		key = uuid.NewV4()
	}

	var expires *time.Time
	if e, ok := options.Context.Value(lifetimeKey).(time.Duration); ok {
		t := time.Now().Add(e)
		expires = aws.Time(t)
	}
	if e, ok := options.Context.Value(expirationKey).(time.Time); ok {
		expires = aws.Time(e)
	}

	metadata := map[string]*string{}
	if m, ok := options.Context.Value(metadataKey).(map[string]*string); ok {
		metadata = m
	}

	var acl *string
	if a, ok := options.Context.Value(aclKey).(string); ok {
		acl = aws.String(a)
	}

	var mime *string
	if m, ok := options.Context.Value(mimetypeKey).(string); ok {
		mime = aws.String(m)
	}

	out, err := s.uploader.Upload(&s3manager.UploadInput{
		Body:        reader,
		ACL:         acl,
		Bucket:      aws.String(s.opts.Bucket),
		Key:         aws.String(key),
		Expires:     expires,
		ContentType: mime,
		Metadata:    metadata,
	})
	if err != nil {
		return "", errors.Wrapf(err, "Failed to upload data to %s/%s\n", s.opts.Bucket, key)
	}
	defer log.Debugf("Successfully created bucket %s and uploaded data with key %s\n", s.opts.Bucket, key)
	return out.Location, nil
}
