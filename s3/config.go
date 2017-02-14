package s3

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type s3Config struct {
	BaseURL string `json:"base_url" config:"store.base_url" default:"http://s3.amazonaws.com/rai-server/"`
	Bucket  string `json:"bucket" config:"store.bucket" default:"rai"`
	ACL     string `json:"acl" config:"store.acl" default:"public-read"`
}

var (
	Config = &s3Config{}
)

func (*s3Config) ConfigName() string {
	return "S3"
}

func (*s3Config) SetDefaults() {
}

func (a *s3Config) Read() {
	vipertags.Fill(a)
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
