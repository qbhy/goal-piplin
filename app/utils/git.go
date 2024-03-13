package utils

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

// CloneRepoBranchOrCommit 克隆指定分支或提交
func CloneRepoBranchOrCommit(repoURL, privateKey, branchOrCommit, destDir string) error {
	auth, err := ssh.NewPublicKeys("git", []byte(privateKey), "")
	if err != nil {
		return fmt.Errorf("error creating ssh auth: %v", err)
	}

	// 首先尝试克隆整个仓库
	r, err := git.PlainClone(destDir, false, &git.CloneOptions{
		URL:  repoURL,
		Auth: auth,
	})
	if err != nil {
		return fmt.Errorf("error cloning repository: %v", err)
	}

	// 获取仓库的工作树
	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("error getting worktree: %v", err)
	}

	// 检出指定的分支或提交
	err = w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(branchOrCommit), // 尝试将 branchOrCommit 视为哈希
	})
	if err != nil {
		// 如果直接使用哈希失败，尝试作为分支名处理
		err = w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.NewBranchReferenceName(branchOrCommit),
		})
		if err != nil {
			return fmt.Errorf("error checking out branch/commit: %v", err)
		}
	}

	return nil
}
