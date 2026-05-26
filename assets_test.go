package webassets

import (
	"io/fs"
	"path"
	"testing"
)

func TestDistFSEmbedsProductionBundle(t *testing.T) {
	t.Run("Should include the SPA index and hashed assets", func(t *testing.T) {
		index, err := DistFS.ReadFile(path.Join(DistDir, "index.html"))
		if err != nil {
			t.Fatalf("DistFS.ReadFile(index.html) error = %v", err)
		}
		if len(index) == 0 {
			t.Fatal("index.html is empty")
		}

		entries, err := fs.ReadDir(DistFS, path.Join(DistDir, "assets"))
		if err != nil {
			t.Fatalf("fs.ReadDir(assets) error = %v", err)
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			switch path.Ext(entry.Name()) {
			case ".js", ".css":
				return
			}
		}
		t.Fatal("expected at least one embedded JavaScript or CSS asset")
	})
}
