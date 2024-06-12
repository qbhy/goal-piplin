package tests

import (
	"fmt"
	"github.com/qbhy/goal-piplin/app/utils"
	"github.com/tj/assert"
	"testing"
)

func TestGitRepoClone(t *testing.T) {
	private := ``
	dir := "/Users/qbhy/project/go/goal-piplin/storage"
	info, err := utils.CloneRepoBranchOrCommit("git@github.com:qbhy/goal-piplin-example.git", private, "master", dir)
	assert.NoError(t, err, err)
	fmt.Println(info)
}

func TestGitRepoCloneWithExec(t *testing.T) {
	private := ``
	dir := "/Users/qbhy/project/go/goal-piplin/storage"
	commit, comment, err := utils.CloneRepo("git@github.com:qbhy/goal-piplin-example.git", private, "master", dir)
	assert.NoError(t, err, err)
	fmt.Println(commit, comment)
}
