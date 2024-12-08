package cache

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/rhoninl/sft/pkg/utils/logger"
)

func pathToCacher(name string) string {
	dir := os.TempDir()
	dir = filepath.Join(dir, "sft")
	return filepath.Join(dir, name)
}

func Cache(name string, data []byte) error {
	if err := os.MkdirAll(pathToCacher(""), 0755); err != nil {
		return err
	}

	logger.Debugf(logger.MoreVerbose, "caching %s", name)
	return os.WriteFile(pathToCacher(name), data, 0644)
}

func Get(name string) ([]byte, error) {
	return os.ReadFile(pathToCacher(name))
}

func GetOrDoAndCache(name string, fetcher func() ([]byte, error)) ([]byte, error) {
	if data, err := Get(name); err == nil {
		logger.Debugf(logger.MoreVerbose, "cache hit: %s, read from cache", name)
		return data, nil
	}

	logger.Debugf(logger.MoreVerbose, "cache miss: %s, fetching from fetcher", name)
	data, err := fetcher()
	if err != nil {
		return nil, err
	}

	if err := Cache(name, data); err != nil {
		return nil, err
	}

	return data, nil
}

func DoWithExpire(name string, expire time.Duration, fetcher func() ([]byte, error)) ([]byte, error) {
	status, err := os.Stat(pathToCacher(name))
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	if err == nil && status.ModTime().Add(expire).After(time.Now()) {
		logger.Debugf(logger.MoreVerbose, "cache hit: %s, read from cache", name)
		return Get(name)
	}

	logger.Debugf(logger.MoreVerbose, "cache miss: %s, fetching from fetcher", name)
	data, err := fetcher()
	if err != nil {
		return nil, err
	}

	if err := Cache(name, data); err != nil {
		return nil, err
	}

	return data, nil
}
