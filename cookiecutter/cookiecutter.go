package cookiecutter

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// CookieCutter ...
type CookieCutter struct {
	rootDir  string
	replacer *strings.Replacer
}

// New creates a CookieCutter
func New(rootDir string, findReplace ...string) (*CookieCutter, error) {
	// validate findReplace
	if len(findReplace)%2 == 1 {
		return nil, errors.New("cookiecutter.New: odd findReplace count")
	}

	// add double curly brace delimiters to find
	for i := 0; i < len(findReplace); i += 2 {
		findReplace[i] = "{{" + findReplace[i] + "}}"
	}

	return &CookieCutter{
		rootDir:  rootDir,
		replacer: strings.NewReplacer(findReplace...),
	}, nil
}

// Run execs the CookieCutter
func (cc *CookieCutter) Run() error {
	return filepath.Walk(cc.rootDir, cc.findReplace)
}

func (cc *CookieCutter) findReplace(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return nil
	}

	read, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	//fmt.Println(string(read))
	// fmt.Println(path)

	newContents := cc.replacer.Replace(string(read))
	// fmt.Println(newContents)

	return ioutil.WriteFile(path, []byte(newContents), 0)
}
