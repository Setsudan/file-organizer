package organizer

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"testing"
	"time"

	"net.ethlny.file-organizer/internal/config"
)

func TestRunOnceMovesFile(t *testing.T) {
	base := t.TempDir()
	src := filepath.Join(base, "Downloads")
	pics := filepath.Join(base, "Pictures")
	others := filepath.Join(base, "Others")
	docs := filepath.Join(base, "Documents")
	music := filepath.Join(base, "Music")
	videos := filepath.Join(base, "Videos")
	apps := filepath.Join(base, "Applications")

	mustMkdir(t, src)
	mustMkdir(t, pics)
	mustMkdir(t, others)
	mustMkdir(t, docs)
	mustMkdir(t, music)
	mustMkdir(t, videos)
	mustMkdir(t, apps)

	input := filepath.Join(src, "photo.png")
	if err := os.WriteFile(input, []byte("data"), 0o644); err != nil {
		t.Fatalf("write input: %v", err)
	}

	cfg := config.Config{
		SourceDir: src,
		Targets: map[string]string{
			"images": pics, "videos": videos, "audio": music, "files": docs,
			"archives": docs, "executable": apps, "code": docs, "others": others,
		},
		MinAge:         0,
		WatchInterval:  time.Second,
		TempExtensions: []string{".crdownload"},
	}
	org, err := New(cfg, slog.New(slog.NewTextHandler(os.Stdout, nil)))
	if err != nil {
		t.Fatalf("new organizer: %v", err)
	}
	if err := org.RunOnce(context.Background()); err != nil {
		t.Fatalf("run once: %v", err)
	}

	out := filepath.Join(pics, "Downloaded Images", "photo.png")
	if _, err := os.Stat(out); err != nil {
		t.Fatalf("expected moved file at %s: %v", out, err)
	}
}

func TestRunOnceSkipsTempExtension(t *testing.T) {
	base := t.TempDir()
	src := filepath.Join(base, "Downloads")
	others := filepath.Join(base, "Others")
	docs := filepath.Join(base, "Documents")
	music := filepath.Join(base, "Music")
	videos := filepath.Join(base, "Videos")
	pics := filepath.Join(base, "Pictures")
	apps := filepath.Join(base, "Applications")

	mustMkdir(t, src)
	mustMkdir(t, others)
	mustMkdir(t, docs)
	mustMkdir(t, music)
	mustMkdir(t, videos)
	mustMkdir(t, pics)
	mustMkdir(t, apps)

	input := filepath.Join(src, "download.crdownload")
	if err := os.WriteFile(input, []byte("partial"), 0o644); err != nil {
		t.Fatalf("write input: %v", err)
	}

	cfg := config.Config{
		SourceDir: src,
		Targets: map[string]string{
			"images": pics, "videos": videos, "audio": music, "files": docs,
			"archives": docs, "executable": apps, "code": docs, "others": others,
		},
		MinAge:         0,
		WatchInterval:  time.Second,
		TempExtensions: []string{".crdownload"},
	}
	org, err := New(cfg, slog.New(slog.NewTextHandler(os.Stdout, nil)))
	if err != nil {
		t.Fatalf("new organizer: %v", err)
	}
	if err := org.RunOnce(context.Background()); err != nil {
		t.Fatalf("run once: %v", err)
	}

	if _, err := os.Stat(input); err != nil {
		t.Fatalf("temp file should remain in source: %v", err)
	}
}

func mustMkdir(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
}
