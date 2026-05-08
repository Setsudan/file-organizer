package organizer

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"net.ethlny.file-organizer/internal/config"
	"net.ethlny.file-organizer/internal/fsops"
)

type Organizer struct {
	cfg      config.Config
	logger   *slog.Logger
	tempExts map[string]struct{}
}

func New(cfg config.Config, logger *slog.Logger) (*Organizer, error) {
	if logger == nil {
		logger = slog.Default()
	}
	if cfg.SourceDir == "" {
		return nil, errors.New("source directory is required")
	}
	if cfg.MinAge < 0 {
		return nil, errors.New("min age cannot be negative")
	}
	if cfg.WatchInterval <= 0 {
		cfg.WatchInterval = 3 * time.Second
	}

	temp := make(map[string]struct{}, len(cfg.TempExtensions))
	for _, ext := range cfg.TempExtensions {
		temp[strings.ToLower(ext)] = struct{}{}
	}

	return &Organizer{cfg: cfg, logger: logger, tempExts: temp}, nil
}

func (o *Organizer) RunOnce(ctx context.Context) error {
	entries, err := os.ReadDir(o.cfg.SourceDir)
	if err != nil {
		return fmt.Errorf("read source dir: %w", err)
	}

	var errs []error
	for _, entry := range entries {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		if err := o.processEntry(entry); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func (o *Organizer) Watch(ctx context.Context) error {
	ticker := time.NewTicker(o.cfg.WatchInterval)
	defer ticker.Stop()

	o.logger.Info("watch mode started", "source", o.cfg.SourceDir, "interval", o.cfg.WatchInterval.String())
	for {
		if err := o.RunOnce(ctx); err != nil && !errors.Is(err, context.Canceled) {
			o.logger.Error("watch pass failed", "error", err.Error())
		}

		select {
		case <-ctx.Done():
			o.logger.Info("watch mode stopped")
			return ctx.Err()
		case <-ticker.C:
		}
	}
}

func (o *Organizer) processEntry(entry os.DirEntry) error {
	if entry.IsDir() {
		return nil
	}

	info, err := entry.Info()
	if err != nil {
		return fmt.Errorf("read file info for %s: %w", entry.Name(), err)
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return nil
	}
	if time.Since(info.ModTime()) < o.cfg.MinAge {
		o.logger.Debug("skipping fresh file", "file", entry.Name())
		return nil
	}

	ext := strings.ToLower(filepath.Ext(entry.Name()))
	if _, ok := o.tempExts[ext]; ok {
		o.logger.Debug("skipping temporary download file", "file", entry.Name())
		return nil
	}

	category := categoryForExtension(ext)
	targetRoot, ok := o.cfg.Targets[string(category)]
	if !ok || targetRoot == "" {
		targetRoot = o.cfg.Targets[string(CategoryOthers)]
	}
	targetDir := filepath.Join(targetRoot, folderNameForCategory(category))
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return fmt.Errorf("create target dir %s: %w", targetDir, err)
	}

	src := filepath.Join(o.cfg.SourceDir, entry.Name())
	dst := filepath.Join(targetDir, entry.Name())
	finalDst, err := fsops.MoveFile(src, dst, fsops.ConflictRename)
	if err != nil {
		return fmt.Errorf("move %s: %w", entry.Name(), err)
	}

	o.logger.Info("file moved", "file", entry.Name(), "category", string(category), "destination", finalDst)
	return nil
}

func folderNameForCategory(category Category) string {
	switch category {
	case CategoryImages:
		return "Downloaded Images"
	case CategoryVideos:
		return "Downloaded Videos"
	case CategoryAudio:
		return "Downloaded Audio"
	case CategoryFiles:
		return "Downloaded Files"
	case CategoryArchives:
		return "Downloaded Archives"
	case CategoryExecutable:
		return "Downloaded Executables"
	case CategoryCode:
		return "Downloaded Code"
	default:
		return "Downloaded Others"
	}
}
