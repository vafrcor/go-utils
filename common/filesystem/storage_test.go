package filesystem

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	assert := assert.New(t)

	// temporary directory
	dir := t.TempDir()

	// file should not exist initially
	p := filepath.Join(dir, "foo.txt")
	assert.False(FileExists(p), "expected %q not to exist", p)

	// create file and verify
	err := os.WriteFile(p, []byte("hello"), 0o644)
	assert.NoError(err, "failed to create file")
	assert.True(FileExists(p), "expected %q to exist", p)

	// remove file and ensure it no longer exists
	err = os.Remove(p)
	assert.NoError(err, "failed to remove file")
	assert.False(FileExists(p), "expected %q not to exist after removal", p)
}

func TestDirExists(t *testing.T) {
	assert := assert.New(t)
	dir := t.TempDir()

	// directory should not exist under base
	sub := filepath.Join(dir, "sub")
	assert.False(DirExists(sub), "expected %q not to exist", sub)

	// create the directory and verify
	err := os.Mkdir(sub, 0o755)
	assert.NoError(err, "failed to create directory")
	assert.True(DirExists(sub), "expected %q to exist", sub)

	// create a file and ensure DirExists returns false
	f := filepath.Join(dir, "file.txt")
	err = os.WriteFile(f, []byte("data"), 0o644)
	assert.NoError(err, "failed to create file")
	assert.False(DirExists(f), "expected DirExists to be false for a file path %q", f)
}

func TestCreatePathDirs(t *testing.T) {
	assert := assert.New(t)
	base := t.TempDir()
	nested := filepath.Join(base, "a", "b", "c.txt")

	// directory does not exist yet
	assert.False(DirExists(filepath.Dir(nested)), "directory should not exist before CreatePathDirs")

	err := CreatePathDirs(nested)
	assert.NoError(err, "CreatePathDirs failed")

	assert.True(DirExists(filepath.Dir(nested)), "expected directory to be created: %q", filepath.Dir(nested))

	// calling again should succeed and leave the dir intact
	err = CreatePathDirs(nested)
	assert.NoError(err, "CreatePathDirs failed on existing path")
}
