package s3

import (
	"net/url"
	"strings"
)

func cleanupKey(key, prefix string) (string, error) {
	keyu, err := url.Parse(key)
	if err != nil {
		log.WithError(err).Error("unable to parse key")
		return "", err
	}

	prefixu, err := url.Parse(prefix)
	if err != nil {
		log.WithError(err).Error("unable to parse prefix")
		return "", err
	}

	return strings.TrimPrefix(keyu.Path, prefixu.Path), nil
}
