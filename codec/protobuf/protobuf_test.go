package protobuf

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vbargl/rainstorm/v6"
	"github.com/vbargl/rainstorm/v6/codec/internal"
)

func TestProtobuf(t *testing.T) {
	u1 := SimpleUser{ID: 1, Name: "John"}
	u2 := SimpleUser{}
	internal.RoundtripTester(t, Codec, &u1, &u2)
	require.True(t, u1.ID == u2.ID)
}

func TestSave(t *testing.T) {
	dir, _ := os.MkdirTemp(os.TempDir(), "rainstorm")
	defer os.RemoveAll(dir)
	db, _ := rainstorm.Open(filepath.Join(dir, "rainstorm.db"), rainstorm.Codec(Codec))
	u1 := SimpleUser{ID: 1, Name: "John"}
	err := db.Save(&u1)
	require.NoError(t, err)
	u2 := SimpleUser{}
	err = db.One("ID", uint64(1), &u2)
	require.NoError(t, err)
	require.Equal(t, u2.Name, u1.Name)
}

func TestGetSet(t *testing.T) {
	dir, _ := os.MkdirTemp(os.TempDir(), "rainstorm")
	defer os.RemoveAll(dir)
	db, _ := rainstorm.Open(filepath.Join(dir, "rainstorm.db"), rainstorm.Codec(Codec))
	err := db.Set("bucket", "key", "value")
	require.NoError(t, err)
	var s string
	err = db.Get("bucket", "key", &s)
	require.NoError(t, err)
	require.Equal(t, "value", s)
}

func TestCodecName(t *testing.T) {
	require.EqualValues(t, Codec.Name(), "protobuf")
}
