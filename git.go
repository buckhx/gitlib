package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	pathlib "path"
	"strings"

	"github.com/buckhx/pathutil"
)

var NOF = []string{}

type Repository struct {
	Path string
}

func (repo *Repository) Op(cmd string, flags []string, args ...string) (out string, err error) {
	flags = append(flags, "-C", repo.Path)
	return Operation(cmd, flags, args...)
}

func (repo *Repository) Init() (err error) {
	_, err = repo.Op("init", NOF)
	return
}

func (repo *Repository) Add(paths []string) (err error) {
	_, err = repo.Op("add", NOF, paths...)
	return
}

func (repo *Repository) Commit(msg string) (err error) {
	_, err = repo.Op("commit", NOF, "-am", msg)
	return
}

// Append these patterns to .git/info/exclude
func (repo *Repository) Exclude(patterns ...string) (err error) {
	excludef, err := os.OpenFile(pathlib.Join(repo.Path, ".git/info/exclude"), os.O_APPEND|os.O_WRONLY, 0600)
	defer excludef.Close()
	if err != nil {
		return
	}
	for _, pattern := range patterns {
		_, err = excludef.WriteString(pattern+"\n")
	}
	return
}

func (repo *Repository) Push() (err error) {
	_, err = repo.Op("push", NOF)
	return
}

func Operation(command string, flags []string, args ...string) (string, error) {
	var stderr, stdout bytes.Buffer
	args = append([]string{command}, args...)
	args = append(flags, args...)
	cmd := exec.Command("git", args...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		err = fmt.Errorf("%s\n%s\n", stderr.String(), strings.Join(append([]string{"git"}, args...), " "))
	}
	return stdout.String(), err
}

func IsRepository(path string) bool {
	return pathutil.Exists(pathlib.Join(path, ".git"))
}

func NewRepository(path string) (repo *Repository, err error) {
	if !pathutil.Exists(path) {
		err = fmt.Errorf("Cannot instantiate Repository because path doesn't exist at %q", path)
	} else {
		repo = &Repository{}
		repo.Path = path
	}
	return
}
