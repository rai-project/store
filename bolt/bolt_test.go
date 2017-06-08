package bolt

import (
	"bytes"
	"testing"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	"github.com/rai-project/store"
	"github.com/stretchr/testify/assert"
)

func TestBoltOperations(t *testing.T) {
	h, err := New(store.Bucket("testing"), BasePath(sourcepath.MustAbsoluteDir()))
	assert.NoError(t, err)
	assert.NotNil(t, h)

	key := "test"
	uploadedString := "test the tester is testing"
	buf := bytes.NewBufferString(uploadedString)
	uploadedKey, err := h.UploadFrom(buf, key)
	assert.NoError(t, err)
	assert.NotEmpty(t, uploadedKey)

	bts, err := h.Get(uploadedKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, bts)

	assert.Equal(t, uploadedString, string(bts))

	defer h.Close()
}
