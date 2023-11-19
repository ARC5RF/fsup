package fsup

import (
	"errors"
	"os"
	"path/filepath"
)

func From(base, what string) (string, error) {
	p := filepath.Join(base, what)
	if _, err := os.Stat(p); !errors.Is(err, os.ErrNotExist) {
		if err != nil {
			return "", err
		}
		return p, nil
	}
	d := filepath.Dir(base)
	if len(d) > 0 && d != base {
		return From(d, what)
	}
	return "", os.ErrNotExist
}

func FromWD(what string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return From(wd, what)
}
