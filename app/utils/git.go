package utils

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
)

// CloneRepoBranchOrCommit 克隆指定分支或提交
func CloneRepoBranchOrCommit(repoURL, publicKey, branchOrCommit, destDir string) error {
	auth, err := ssh.NewPublicKeys("git", []byte(publicKey), "")
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

// GetRepositoryBranchesAndTags 获取 Git 仓库的分支和 Tags
func GetRepositoryBranchesAndTags(repoURL string, publicKey string) ([]string, []string, error) {
	auth, err := ssh.NewPublicKeys("git", []byte(publicKey), "")
	if err != nil {
		return nil, nil, fmt.Errorf("error creating ssh auth: %v", err)
	}

	// 克隆仓库到内存中
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:  repoURL,
		Auth: auth,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("unable to clone repository: %w", err)
	}

	// 获取所有分支
	branches, err := repo.Branches()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to list branches: %w", err)
	}

	var branchList []string
	err = branches.ForEach(func(ref *plumbing.Reference) error {
		branchList = append(branchList, ref.Name().Short())
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	// 获取所有 Tags
	tags, err := repo.Tags()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to list tags: %w", err)
	}

	var tagList []string
	err = tags.ForEach(func(ref *plumbing.Reference) error {
		tagList = append(tagList, ref.Name().Short())
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return branchList, tagList, nil
}
