package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rai-project/aws"
	"github.com/rai-project/config"
	"github.com/rai-project/store"
)

type options struct {
}

type s3Client struct {
	client     *s3.S3
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
	s3Opts     options
	opts       *store.Options
}

func New(iopts ...store.Option) (store.Store, error) {
	s3Opts := options{}
	opts := NewOptions()

	for _, o := range iopts {
		o(opts)
	}

	var sess *session.Session
	if s, ok := opts.Context.Value(sessionKey{}).(*session.Session); ok && s != nil {
		sess = s.Copy()
	}
	if sess == nil {
		var err error
		sess, err = aws.NewSession()
		if err != nil {
			return nil, err
		}
	}

	sess.Config.WithEndpoint(opts.BaseURL)

	if config.IsVerbose || config.IsDebug {
		sess.Config.WithCredentialsChainVerboseErrors(true).WithLogger(log)
	}

	client := s3.New(sess)
	uploader := s3manager.NewUploaderWithClient(client)
	downloader := s3manager.NewDownloaderWithClient(client)

	return &s3Client{
		client:     client,
		uploader:   uploader,
		downloader: downloader,
		s3Opts:     s3Opts,
		opts:       opts,
	}, nil
}

func (s *s3Client) Options() store.Options {
	return *s.opts
}

func (*s3Client) Name() string {
	return "S3"
}

func (c *s3Client) Close() error {
	return nil
}
