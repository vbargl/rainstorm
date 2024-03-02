package json

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vbargl/rainstorm/v6/codec/internal"
)

func TestJSON(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}

func TestCodecName(t *testing.T) {
	require.EqualValues(t, Codec.Name(), "json")
}
