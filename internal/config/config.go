package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	SourceDir      string            `json:"sourceDir"`
	Targets        map[string]string `json:"targets"`
	MinAge         time.Duration     `json:"-"`
	WatchInterval  time.Duration     `json:"-"`
	TempExtensions []string          `json:"tempExtensions"`
}

type fileConfig struct {
	SourceDir      string            `json:"sourceDir"`
	Targets        map[string]string `json:"targets"`
	TempExtensions []string          `json:"tempExtensions"`
}

func Load(path string) (Config, error) {
	cfg, err := defaultConfig()
	if err != nil {
		return Config{}, err
	}
	if path == "" {
		return cfg, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}

	var userCfg fileConfig
	if err := json.Unmarshal(data, &userCfg); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}

	if userCfg.SourceDir != "" {
		cfg.SourceDir = filepath.Clean(userCfg.SourceDir)
	}
	for k, v := range userCfg.Targets {
		if v != "" {
			cfg.Targets[k] = filepath.Clean(v)
		}
	}
	if len(userCfg.TempExtensions) > 0 {
		cfg.TempExtensions = append([]string(nil), userCfg.TempExtensions...)
	}
	return cfg, nil
}

func defaultConfig() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("resolve home dir: %w", err)
	}

	return Config{
		SourceDir: filepath.Join(homeDir, "Downloads"),
		Targets: map[string]string{
			"images":     filepath.Join(homeDir, "Downloads"),
			"videos":     filepath.Join(homeDir, "Downloads"),
			"audio":      filepath.Join(homeDir, "Downloads"),
			"files":      filepath.Join(homeDir, "Downloads"),
			"archives":   filepath.Join(homeDir, "Downloads"),
			"executable": filepath.Join(homeDir, "Downloads"),
			"code":       filepath.Join(homeDir, "Downloads"),
			"others":     filepath.Join(homeDir, "Downloads"),
		},
		MinAge:        5 * time.Second,
		WatchInterval: 3 * time.Second,
		TempExtensions: []string{
			".crdownload", ".part", ".tmp", ".opdownload", ".download",
		},
	}, nil
}
