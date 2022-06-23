package ethash

import (
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func TestEnsureSize(t *testing.T) {
	t.Run("epoch=2", testEnsureSize(cacheSize(2)))
	t.Run("epoch=20", testEnsureSize(cacheSize(20)))
	t.Run("epoch=200", testEnsureSize(cacheSize(200)))
	t.Run("epoch=2000", testEnsureSize(cacheSize(2000)))
}

func testEnsureSize(size uint64) func(t *testing.T) {
	return func(t *testing.T) {
		path := filepath.Join(os.TempDir(), "geth-fallocate-test-"+strconv.Itoa(rand.Int()))
		defer os.RemoveAll(path)

		// Ensure the data folder exists
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		// Create a huge temporary empty file to fill with data
		temp := path + "." + strconv.Itoa(rand.Int())

		dump, err := os.Create(temp)
		if err != nil {
			t.Fatal(err)
		}
		if err = ensureSize(dump, int64(len(dumpMagic))*4+int64(size)); err != nil {
			dump.Close()
			os.Remove(temp)
			t.Fatal(err)
		}
	}
}
