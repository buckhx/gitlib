package git

import (
	"fmt"
	"os/exec"
	"bytes"
	pathlib "path"

	"github.com/buckhx/pathutil"
)

var NOF := []string{""}

type Repository struct {
	Path string
}

func (repo *Repository) Op(cmd string, flags []string, ...args string) (string, error) {
	flags = append(flags, "-C", repo.Path)
	return Op(cmd, flags, args...)
}

func (repo *Repository) Init() (err error) {
	_, err := repo.Op("init", NOF)
	return
}

func Add(paths []string) (err error) {
	_, err := repo.Op("add", NOF, paths...)
	return
}

func Commit(msg string) (err error) {
	_, err := repo.Op("commit", NOF, "-am", msg)
	return
	
}

func Push() (err error) {
	_, err := repo.Op("push", NOF)
	return
}

func Op(cmd string, flags []string, ...args string) (string, error) {
	var stderr, stdout bytes.Buffer	
	args = append([]string{cmd}, args...)
	args = append(flags, args...)
	cmd := exec.Command("git", args...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		err = fmt.Errorf("%s\n%s\n", stderr.String(), strings.Join(append([]string{"git"}, args...)," "))
	}
	return stdout.String(), err
}

func IsRepository(path string) bool {
	return pathutil.Exists(pathlib.Join(path, ".git"))
}

func NewRepository(path string) (repo Repository, err error) {
	if !pathutil.Exists(path) {
		err := fmt.Errorf("Cannot instantiate Repository because path doesn't exist at %q", path)
		repo := nil
	} else {
		repo := &Repository
		repo.Path = path
		err := nil
	}
	return 
}
