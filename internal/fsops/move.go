package fsops

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ConflictPolicy string

const (
	ConflictRename ConflictPolicy = "rename"
)

func MoveFile(src, dst string, policy ConflictPolicy) (string, error) {
	finalDst, err := resolveConflict(dst, policy)
	if err != nil {
		return "", err
	}

	if err := os.Rename(src, finalDst); err == nil {
		return finalDst, nil
	} else if !isCrossDeviceError(err) {
		return "", fmt.Errorf("rename file: %w", err)
	}

	if err := copyFile(src, finalDst); err != nil {
		return "", fmt.Errorf("copy fallback: %w", err)
	}
	if err := os.Remove(src); err != nil {
		return "", fmt.Errorf("cleanup source after copy: %w", err)
	}
	return finalDst, nil
}

func resolveConflict(dst string, policy ConflictPolicy) (string, error) {
	if policy != ConflictRename {
		return "", fmt.Errorf("unsupported conflict policy: %s", policy)
	}
	if _, err := os.Stat(dst); errors.Is(err, os.ErrNotExist) {
		return dst, nil
	} else if err != nil {
		return "", fmt.Errorf("check destination: %w", err)
	}

	ext := filepath.Ext(dst)
	base := strings.TrimSuffix(filepath.Base(dst), ext)
	dir := filepath.Dir(dst)

	for i := 1; i <= 10000; i++ {
		candidate := filepath.Join(dir, fmt.Sprintf("%s (%d)%s", base, i, ext))
		if _, err := os.Stat(candidate); errors.Is(err, os.ErrNotExist) {
			return candidate, nil
		} else if err != nil {
			return "", fmt.Errorf("check destination candidate: %w", err)
		}
	}
	return "", fmt.Errorf("too many name conflicts for %s", dst)
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
}

func isCrossDeviceError(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "cross-device")
}
