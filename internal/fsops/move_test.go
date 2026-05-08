package fsops

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMoveFileRenamesOnConflict(t *testing.T) {
	base := t.TempDir()
	srcDir := filepath.Join(base, "src")
	dstDir := filepath.Join(base, "dst")
	if err := os.MkdirAll(srcDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		t.Fatal(err)
	}

	src := filepath.Join(srcDir, "file.txt")
	dst := filepath.Join(dstDir, "file.txt")
	if err := os.WriteFile(src, []byte("new"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(dst, []byte("old"), 0o644); err != nil {
		t.Fatal(err)
	}

	finalDst, err := MoveFile(src, dst, ConflictRename)
	if err != nil {
		t.Fatalf("move file: %v", err)
	}
	if finalDst == dst {
		t.Fatalf("expected conflict rename destination, got %s", finalDst)
	}
	if _, err := os.Stat(finalDst); err != nil {
		t.Fatalf("moved file missing: %v", err)
	}
}
