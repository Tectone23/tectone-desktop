package git

import (
	"errors"
	"io"

	"github.com/go-git/go-git/v5"
)

type Commit struct {
	Hash string
}

func Pull(path string) (*Commit, error) {
	var c *Commit

	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(path)
	if err != nil {
		return c, err
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return c, err
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	if err := w.Pull(&git.PullOptions{RemoteName: "origin"}); err != nil {
		return c, err
	}

	// Print the latest commit that was just pulled
	ref, err := r.Head()
	if err != nil {
		return c, err
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return c, err
	}

	c = &Commit{
		Hash: commit.Hash.String(),
	}

	return c, nil
}

func Clone(path, url string, w1, w2 io.Writer) error {
	repo, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          w1,
	})

	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}

	if err := repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		Progress:   w2,
	}); err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return err
	}

	return nil
}

func Fetch(path, url string, w io.Writer) error {
	return nil
}
