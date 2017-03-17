package s3

import (
	"strings"

	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type s3Config struct {
	Provider string        `json:"provider" config:"store.provider" default:"s3"`
	BaseURL  string        `json:"base_url" config:"store.base_url" default:"http://s3.amazonaws.com/rai-server/"`
	Bucket   string        `json:"bucket" config:"store.bucket" default:"rai"`
	ACL      string        `json:"acl" config:"store.acl" default:"public-read"`
	done     chan struct{} `json:"-" config:"-"`
}

var (
	Config = &s3Config{
		done: make(chan struct{}),
	}
)

func (*s3Config) ConfigName() string {
	return "S3"
}

func (a *s3Config) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *s3Config) Read() {
	vipertags.Fill(a)
	if !strings.HasPrefix(a.BaseURL, "http://") && !strings.HasPrefix(a.BaseURL, "https://") {
		a.BaseURL = "http://" + a.BaseURL
	}
}

func (c s3Config) Wait() {
	<-c.done
}

func (c *s3Config) String() string {
	return pp.Sprintln(c)
}

func (c *s3Config) Debug() {
	log.Debug("S3 Config = ", c)
}

func init() {
	config.Register(Config)
}
