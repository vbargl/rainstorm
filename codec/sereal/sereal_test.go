package sereal

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vbargl/rainstorm/v6/codec/internal"
)

type SerealUser struct {
	Name string
	Self *SerealUser
}

func TestSereal(t *testing.T) {
	u1 := &SerealUser{Name: "Sereal"}
	u1.Self = u1 // cyclic ref
	u2 := &SerealUser{}
	internal.RoundtripTester(t, Codec, &u1, &u2)
	require.True(t, u2 == u2.Self)
}

func TestCodecName(t *testing.T) {
	require.EqualValues(t, Codec.Name(), "sereal")
}
