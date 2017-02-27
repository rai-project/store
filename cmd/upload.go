package cmd

import (
	"io"
	"os"
	"time"

	"github.com/Unknwon/com"
	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
	"github.com/rai-project/archive"
	"github.com/rai-project/store"
	"github.com/rai-project/store/s3"
	"github.com/rai-project/uuid"
	"github.com/spf13/cobra"
)

var uploadKey string

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use: "upload",
	Aliases: []string{
		"put",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("expected a file or directory to upload")
		}
		fileName := args[0]
		if !com.IsFile(fileName) && !com.IsDir(fileName) {
			return errors.Errorf("file or directory %v was not found", fileName)
		}

		var err error
		var reader io.ReadCloser

		if com.IsDir(fileName) {
			reader, err = archive.Zip(fileName)
			if err != nil {
				return errors.Wrapf(err, "unable to archive %v", fileName)
			}
		} else {
			reader, err = os.Open(fileName)
			if err != nil {
				return errors.Wrapf(err, "unable to open file %v", fileName)
			}
		}

		defer reader.Close()

		if uploadKey == "" {
			uploadKey = uuid.New(fileName)
		}

		str, err := s3.New(
			store.BaseURL("http://s3.amazonaws.com/rai-server/"),
		)
		if err != nil {
			return errors.Wrapf(err, "unable to create an s3 connection")
		}

		key, err := str.UploadFrom(reader,
			uploadKey,
			s3.Expiration(time.Now().AddDate(0, 1, 0)),
		)
		if err != nil {
			return err
		}

		pp.Println(key)

		return nil
	},
}

func init() {
	uploadCmd.PersistentFlags().StringVarP(&uploadKey, "key", "k", "", "upload key")
}
