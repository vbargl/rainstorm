package aes

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vbargl/rainstorm/v6/codec/internal"
	"github.com/vbargl/rainstorm/v6/codec/json"
)

var testKey, _ = base64.StdEncoding.DecodeString("xkBTXc1wn0C/aL31u9SA7g==")

func TestAES(t *testing.T) {
	aes, err := NewAES(json.Codec, testKey)
	require.NoError(t, err)

	internal.RoundtripTester(t, aes)
}
