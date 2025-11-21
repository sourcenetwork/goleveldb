//go:build js
// +build js

package storage

import (
	"os"
)

type jsFileLock struct{}

func (fl *jsFileLock) release() error {
	return nil
}

func newFileLock(path string, readOnly bool) (fl fileLock, err error) {
	return &jsFileLock{}, nil
}

func rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func syncDir(name string) error {
	return nil
}
