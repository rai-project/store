package s3

import (
	"context"
	"io"
	"mime"
	"os"
	"runtime"
	"time"

	"path/filepath"

	"bytes"

	"github.com/Unknwon/com"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	humanize "github.com/dustin/go-humanize"
	"github.com/pkg/errors"
	"github.com/rai-project/store"
	"github.com/rai-project/uuid"
	"gopkg.in/cheggaaa/pb.v1"
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

func newProgress(output io.Writer, bytes int64) *pb.ProgressBar {
	// get the new original progress bar.
	bar := pb.New64(bytes)

	// Set new human friendly print units.
	bar.SetUnits(pb.U_BYTES)

	// Set output to be stdout
	bar.Output = output

	// Show current speed is true.
	bar.ShowSpeed = true

	// Refresh rate for progress bar is set to 100 milliseconds.
	bar.SetRefreshRate(time.Millisecond * 100)

	// Use different unicodes for Linux, OS X and Windows.
	switch runtime.GOOS {
	case "linux":
		// Need to add '\x00' as delimiter for unicode characters.
		bar.Format("┃\x00▓\x00█\x00░\x00┃")
	case "darwin":
		// Need to add '\x00' as delimiter for unicode characters.
		bar.Format(" \x00▓\x00 \x00░\x00 ")
	default:
		// Default to non unicode characters.
		bar.Format("[=> ]")
	}
	return bar
}

func (s *s3Client) Upload(path string, key string, opts ...store.UploadOption) (string, error) {

	if !com.IsFile(path) {
		return "", errors.Errorf("The file %v was not found", path)
	}

	ext := filepath.Ext(path)
	contentType := mime.TypeByExtension(ext)

	options := store.UploadOptions{
		Context: context.WithValue(
			context.Background(),
			contentTypeKey{},
			contentType,
		),
	}

	for _, o := range opts {
		o(&options)
	}

	if fileSizeLimit, ok := options.Context.Value(fileSizeLimitKey{}).(int64); ok {
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

	if options.ProgressOutput != nil && options.Progress == nil {
		stats, err := file.Stat()
		if err == nil {
			opts = append(
				opts,
				store.UploadProgress(newProgress(options.ProgressOutput, stats.Size())),
			)
		}
	}

	return s.UploadFrom(file, key, opts...)
}

func (s *s3Client) UploadFrom(reader io.Reader, key string, opts ...store.UploadOption) (string, error) {

	options := store.UploadOptions{
		Context: context.WithValue(
			context.Background(),
			aclKey{},
			Config.ACL,
		),
		Metadata: map[string]string{},
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
	if e, ok := options.Context.Value(lifetimeKey{}).(time.Duration); ok {
		t := time.Now().Add(e)
		expires = aws.Time(t)
	}
	if e, ok := options.Context.Value(expirationKey{}).(time.Time); ok {
		expires = aws.Time(e)
	}

	metadata := map[string]*string{}
	for k, v := range options.Metadata {
		metadata[k] = aws.String(v)
	}

	var acl *string
	if a, ok := options.Context.Value(aclKey{}).(string); ok {
		acl = aws.String(a)
	}

	var contentType *string
	if m, ok := options.Context.Value(contentTypeKey{}).(string); ok {
		contentType = aws.String(m)
	}

	progress := options.Progress
	if options.ProgressOutput != nil && progress == nil {
		buf := new(bytes.Buffer)
		size, err := buf.ReadFrom(reader)
		if err != nil {
			return "", errors.Wrap(err, "cannot determine size of reader")
		}
		progress = newProgress(options.ProgressOutput, size)
		reader = progress.NewProxyReader(buf)
	} else if progress != nil {
		reader = progress.NewProxyReader(reader)
	}

	if progress != nil {
		progress.Start()
	}

	out, err := s.uploader.UploadWithContext(
		options.Context,
		&s3manager.UploadInput{
			Body:        reader,
			ACL:         acl,
			Bucket:      aws.String(s.opts.Bucket),
			Key:         aws.String(key),
			Expires:     expires,
			ContentType: contentType,
			Metadata:    metadata,
		},
	)
	if err != nil {
		return "", errors.Wrapf(err, "Failed to upload data to %s/%s\n", s.opts.Bucket, key)
	}

	if progress != nil {
		if options.ProgressFinishMessage == "" {
			progress.Finish()
		} else {
			progress.FinishPrint(options.ProgressFinishMessage)
		}
	}

	defer log.Debugf("Successfully created bucket %s and uploaded data with key %s\n", s.opts.Bucket, key)
	return out.Location, nil
}
