package msgpack

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vbargl/rainstorm/v6/codec/internal"
)

func TestMsgpack(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}

func TestCodecName(t *testing.T) {
	require.EqualValues(t, Codec.Name(), "msgpack")
}
