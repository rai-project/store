package s3

import "strings"

func cleanupKey(key, prefix string) string {
	return strings.TrimPrefix(
		strings.TrimPrefix(
			strings.TrimPrefix(
				key,
				"http://",
			),
			"https://",
		),
		prefix,
	)
}
