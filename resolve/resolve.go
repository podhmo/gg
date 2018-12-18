package resolve

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

var (
	// ErrNotFound is the error that package is not fount in current go env
	ErrNotFound = fmt.Errorf("not found")
)

// Main :
func Main(args []string) error {
	c := &Config{
		OnError: func(path string, err error) {
			log.Println(path) // xxx
		},
		OnResult: func(filepath string, pkgpath string) {
			fmt.Println(pkgpath)
		},
	}
	return c.Run(args)
}

// Config :
type Config struct {
	OnError  func(filepath string, err error)
	OnResult func(filepath string, pkgpath string)
}

// Run :
func (c *Config) Run(args []string) error {
	for _, path := range args {
		pkgpath, err := resolve(path)
		if err != nil {
			c.OnError(path, err)
			continue
		}
		c.OnResult(path, pkgpath)
	}
	return nil
}

func resolve(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		u, err := user.Current()
		if err != nil {
			return "", err
		}
		path = filepath.Join(u.HomeDir, path[1:])
	} else if path == "." || path == "./" {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		path = wd
	} else if strings.HasPrefix(path, "./") {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		path = strings.Replace(path, ".", wd, 1)
	} else if strings.HasPrefix(path, "../") {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		path = wd + "/" + path
	}

	ctxt := build.Default
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	for _, srcdir := range ctxt.SrcDirs() {
		if strings.HasPrefix(path, srcdir) {
			pkgpath := strings.TrimLeft(strings.Replace(path, srcdir, "", 1), "/")
			return pkgpath, nil
		}
	}
	return "", errors.Wrapf(ErrNotFound, "%q is not subdir of srcdirs(%q)", path, build.Default.SrcDirs())
}
