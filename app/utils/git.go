package utils

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
	"os"
	"os/exec"
)

// CloneRepoBranchOrCommit 克隆指定分支或提交
func CloneRepoBranchOrCommit(repoURL, publicKey, branchOrCommit, destDir string) error {
	auth, err := ssh.NewPublicKeys("git", []byte(publicKey), "")
	if err != nil {
		return fmt.Errorf("error creating ssh auth: %v", err)
	}

	var isBranch bool
	references := []plumbing.ReferenceName{
		plumbing.NewBranchReferenceName(branchOrCommit),
		plumbing.NewTagReferenceName(branchOrCommit),
		"",
	}

	var registry *git.Repository

	_ = os.RemoveAll(destDir)

	for i, reference := range references {
		// 首先尝试克隆整个仓库
		registry, err = git.PlainClone(destDir, false, &git.CloneOptions{
			URL:           repoURL,
			Auth:          auth,
			ReferenceName: reference,
		})

		if err == nil {
			if i < len(references)-1 {
				isBranch = true
			}
			break
		} else if i == len(references)-1 {
			return err
		}
	}

	if err != nil || registry == nil {
		return fmt.Errorf("error cloning repository: %v", err)
	}

	if !isBranch {
		// Change to the specified directory
		if err = os.Chdir(destDir); err != nil {
			return fmt.Errorf("failed to change directory: %v", err)
		}

		// Run 'git checkout' command
		cmd := exec.Command("git", "checkout", branchOrCommit)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err = cmd.Run(); err != nil {
			return fmt.Errorf("failed to execute 'git checkout': %v", err)
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
